package list

//GetIndex returns the Item of index i in the List.
func (l *List) GetIndex(i uint64) *Item {
	if i >= l.length {
		panic(ErrIndexOutOfRange)
	}
	if i < l.length/2 {
		//Start from the begining if we're in the first half of the list
		item := l.start
		for ; i > 0 && item != nil; i-- {
			item = item.after
		}
		return item
	}
	//Start from the end if we're in the second half of the list
	item := l.end
	i = l.length - 1 - i
	for ; i > 0 && item != nil; i-- {
		item = item.before
	}
	return item
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
func (i *Item) Forward(n uint) *Item {
	for ; n > 0 && i != nil; n-- {
		i = i.after
	}
	return i
}

//Backward returns the Item n Items backwards in the list or nil if you leave the list.
func (i *Item) Backward(n uint) *Item {
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
