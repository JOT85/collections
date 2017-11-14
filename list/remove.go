package list

func (l *List) remove(i *Item) {
	//There is now 1 less item
	l.length--
	first := l.isFirstItem(i)
	last := l.isLastItem(i)
	if last && first {
		//If this is the only item, only set the start and end of the list to nil
		l.end = nil
		l.start = nil
	} else if last {
		//If we are only the last item, set the end of the list to the Item before this, set the Item after the new end to nil, and set the Item before this to nil
		l.end = i.before
		i.before.after = nil
		i.before = nil
	} else if first {
		//If we are only the first item, set the start of the list to the Item after this, set the Item before the new start to nil, and set the Item after this to nil
		l.start = i.after
		i.after.before = nil
		i.after = nil
	} else {
		//If we are in the middle of the list, set the Item after the Item before this to the current Item after this, the Item before the Item after this to the current Item before this, and remove links from this Item
		i.before.after = i.after
		i.after.before = i.before
		i.before = nil
		i.after = nil
	}
}

//Remove removes the given Item from the List. It DOES NOT check that the item is from this List, so it can cause some unexpected behavior if you remove an Item from another List.
func (l *List) Remove(i *Item) {
	//Check this Item is part of a List
	l.PanicIfNotInList(i)
	l.PanicIfEmpty()
	l.remove(i)
}

//Pop removes and returns the last Item in the List
func (l *List) Pop() *Item {
	i := l.end
	l.remove(i)
	return i
}

//Shift removes and returns the first Item in the List
func (l *List) Shift() *Item {
	i := l.start
	l.remove(i)
	return i
}
