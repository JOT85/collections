package list

//ItemPosition return position information about the given Item in the List. It returns a bool, true if it is in the List, false if not.
//And returns a uint64 of the index of the Item, or if it isn't in the list, it is the length of the List.
func (l *List) ItemPosition(search *Item, startIndex uint64, searchDown, circle bool) (bool, uint64) {
	//If the list has no Items, then it won't be in it
	if l.length == 0 {
		return false, 0
	}
	//Make sure the Item is linked
	l.PanicIfNotInList(search)
	//Start at the startIndex
	i := startIndex
	item := l.GetIndex(startIndex)
	//Stop at 0 if we are searching down, or the length if we are searching up
	var stopAt uint64
	if !searchDown {
		stopAt = l.length
		//Don't circle if we don't need to
		if startIndex == 0 {
			circle = false
		}
	} else if startIndex == l.length {
		//Don't circle if we don't need to
		circle = false
	}
	for {
		//If this is the Item we are searching for, return the index
		if item == search {
			return true, i
		}
		//Move on to the next Item
		if searchDown {
			item = item.before
		} else {
			item = item.after
		}
		//If this is the stopping point, break if we don't need to circle, do circle logic if we do
		if i == stopAt {
			if circle {
				if searchDown {
					//If we are searching down, then we now need to search from the last item to the one after where we started
					i = l.length - 1
					stopAt = startIndex + 1
				} else {
					//If we are searching up, then we need to search from 0 to one before where we started
					i = 0
					stopAt = startIndex - 1
				}
				//And we don't need to circle again
				circle = false
				continue
			}
			break
		}
		//Change the index
		if searchDown {
			i--
		} else {
			i++
		}
	}
	return false, i
}

//IndexOf searches all the values in list from the startIndex (counting up if searchDown is false) and compares the value to the given search.
//If the value is equal to the search, it returns true and the index of the current Item.
//If it is not found, it returns false and the last checked index.
func (l *List) IndexOf(search interface{}, startIndex uint64, searchDown, circle bool) (bool, uint64) {
	//If the list has no Items, then it won't be in it
	if l.length == 0 {
		return false, 0
	}
	//Start at the startIndex
	i := startIndex
	item := l.GetIndex(startIndex)
	//Stop at 0 if we are searching down, or the length if we are searching up
	var stopAt uint64
	if !searchDown {
		stopAt = l.length
		//Don't circle if we don't need to
		if startIndex == 0 {
			circle = false
		}
	} else if startIndex == l.length {
		//Don't circle if we don't need to
		circle = false
	}
	for {
		//If this is the Item we are searching for, return the index
		if *(*interface{})(item.value) == search {
			return true, i
		}
		//Move on to the next Item
		if searchDown {
			item = item.before
		} else {
			item = item.after
		}
		//If this is the stopping point, break if we don't need to circle, do circle logic if we do
		if i == stopAt {
			if circle {
				if searchDown {
					//If we are searching down, then we now need to search from the last item to the one after where we started
					i = l.length - 1
					stopAt = startIndex + 1
				} else {
					//If we are searching up, then we need to search from 0 to one before where we started
					i = 0
					stopAt = startIndex - 1
				}
				//And we don't need to circle again
				circle = false
				continue
			}
			break
		}
		//Change the index
		if searchDown {
			i--
		} else {
			i++
		}
	}
	return false, i
}
