package list

//GetIndex returns the Item of index i in the List.
func (l *List) GetIndex(i uint64) *Item {
	if i >= l.length {
		panic(ErrIndexOutOfRange)
	}
	//Return if it is the start or the end
	if i == 0 {
		return l.start
	}
	if i == l.length-1 {
		return l.end
	}
	//Or if it was the last thing we indexed
	if i == l.lastIndexedIndex {
		if l.lastIndexedItem != nil {
			return l.lastIndexedItem
		}
	}
	//If we're between the last indexed and the end
	if i > l.lastIndexedIndex {
		distFromLastIndexed := i - l.lastIndexedIndex
		distFromEnd := l.length - 1 - i
		//Go the shortest way
		if distFromLastIndexed > distFromEnd {
			return l.end.Backward(distFromEnd)
		}
		return l.lastIndexedItem.Forward(distFromLastIndexed)
	}
	distFromLastIndexed := l.lastIndexedIndex - i
	//Go the shortest way
	if distFromLastIndexed > i {
		return l.start.Forward(i)
	}
	return l.lastIndexedItem.Backward(distFromLastIndexed)
}

//Previous returns the previous item in the List or nil if there isn't one.
func (i *Item) Previous() *Item {
	return i.before
}

//Next returns the next item in the List or nil if there isn't one.
func (i *Item) Next() *Item {
	return i.after
}

//Forward returns the Item n Items forward in the list, or nil if you leave the list.
func (i *Item) Forward(n uint64) *Item {
	for ; n > 0 && i != nil; n-- {
		i = i.after
	}
	return i
}

//Backward returns the Item n Items backwards in the list or nil if you leave the list.
func (i *Item) Backward(n uint64) *Item {
	for ; n > 0 && i != nil; n-- {
		i = i.before
	}
	return i
}

//IsEmpty returns true if there are no items in queue, false if otherwise
func (l *List) IsEmpty() bool {
	return l.start == nil
}

//Length returns the amount of items in the queue
func (l *List) Length() (n uint64) {
	return l.length
}
