package list

func (l *List) isFirstItem(i *Item) bool {
	if i.before != nil {
		return false
	}
	if l.start == i {
		return true
	}
	panic(ErrNotInList)
}

func (l *List) isLastItem(i *Item) bool {
	if i.after != nil {
		return false
	}
	if l.end == i {
		return true
	}
	panic(ErrNotInList)
}

//PanicIfEmpty panics with ErrNoItems if the List is empty or incomplete.
func (l *List) PanicIfEmpty() {
	//Check this List has some Items
	if l.length == 0 || l.start == nil || l.end == nil {
		panic(ErrNoItems)
	}
}

//PanicIfNotInList panics with ErrNotInList if i cannot be in this List, it DOES NOT check it is in the List (see ), just simple checks to see if it is possible.
func (l *List) PanicIfNotInList(i *Item) {
	if (i.before == nil && i.after == nil && l.length != 1) || l.length == 0 {
		panic(ErrNotInList)
	}
}

//PanicIfNotInListCheck panics with ErrNotInList if i is not in the List. It does iterate through to check so isn't very efficient, in most cases, PanicIfNotInList may be sufficient - however it won't prevent really stupid logic errors.
func (l *List) PanicIfNotInListCheck(i *Item) {
	inList, _ := l.ItemPosition(i, 0, false, false)
	if !inList {
		panic(ErrNotInList)
	}
}

//PanicIfInList panics with ErrInList if the Item is in a List
func (i *Item) PanicIfInList() {
	if i.before != nil || i.after != nil {
		panic(ErrInList)
	}
}
