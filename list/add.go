package list

import "unsafe"

//insertAssumingEmpty inserts the Item at index 0 assuming there are no Items in the List.
func (l *List) insertAssumingEmpty(insert *Item) {
	//The length is now 1, and this Item is the start and the end
	l.length = 1
	l.start = insert
	l.end = insert
}

//insertBefore inserts the second Item before ther first Item in the List.
func (l *List) insertBefore(item, insert *Item) {
	l.length++
	if l.isFirstItem(item) {
		l.start = insert
	} else {
		item.before.after = insert
	}
	insert.before = item.before
	insert.after = item
	item.before = insert
}

//insertAfter inserts the second Item after ther first Item in the List.
func (l *List) insertAfter(item, insert *Item) {
	l.length++
	if l.isLastItem(item) {
		l.end = insert
	} else {
		item.after.before = insert
	}
	insert.before = item
	insert.after = item.after
	item.after = insert
}

//InsertItemBefore inserts the given item (must not be part of a list) before the given Item in this List.
func (l *List) InsertItemBefore(item, insert *Item) {
	insert.PanicIfInList()
	l.PanicIfNotInList(item)
	l.insertBefore(item, insert)
}

//InsertItemAfter inserts the given item (must not be part of a list) after the given Item in this List.
func (l *List) InsertItemAfter(item, insert *Item) {
	insert.PanicIfInList()
	l.PanicIfNotInList(item)
	l.insertAfter(item, insert)
}

//InsertPointerBefore inserts a new item, holding the value given, before the Item. And returns a pointer to the new Item.
func (l *List) InsertPointerBefore(item *Item, value unsafe.Pointer) *Item {
	ni := &Item{
		value,
		nil, nil,
	}
	l.insertBefore(item, ni)
	return ni
}

//InsertPointerAfter inserts a new item, holding the value given, after the Item. And returns a pointer to the new Item.
func (l *List) InsertPointerAfter(item *Item, value unsafe.Pointer) *Item {
	ni := &Item{
		value,
		nil, nil,
	}
	l.insertAfter(item, ni)
	return ni
}

//InsertBefore wraps InsertPointerBefore by calling InsertPointerBefore with a pointer to the given value.
func (l *List) InsertBefore(item *Item, value interface{}) *Item {
	return l.InsertPointerBefore(item, unsafe.Pointer(&value))
}

//InsertAfter wraps InsertPointerAfter by calling InsertPointerAfter with a pointer to the given value.
func (l *List) InsertAfter(item *Item, value interface{}) *Item {
	return l.InsertPointerAfter(item, unsafe.Pointer(&value))
}

//UnshiftItem inserts the Item at the start of the list.
func (l *List) UnshiftItem(item *Item) {
	if l.length == 0 {
		l.insertAssumingEmpty(item)
	} else {
		l.insertBefore(l.start, item)
	}
}

//PushItem inserts the Item at the end of the list.
func (l *List) PushItem(item *Item) {
	if l.length == 0 {
		l.insertAssumingEmpty(item)
	} else {
		l.insertAfter(l.end, item)
	}
}

//UnshiftPointer creates a new Item containing the pointer and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) UnshiftPointer(value unsafe.Pointer) *Item {
	item := &Item{
		value,
		nil, nil,
	}
	l.UnshiftItem(item)
	return item
}

//PushPointer creates a new Item containing the pointer and calls PushItem with the new Item. Returns the new Item.
func (l *List) PushPointer(value unsafe.Pointer) *Item {
	item := &Item{
		value,
		nil, nil,
	}
	l.PushItem(item)
	return item
}

//Unshift creates a new Item containing the given value and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) Unshift(value interface{}) *Item {
	return l.UnshiftPointer(unsafe.Pointer(&value))
}

//Push creates a new Item containing the given value and calls PushItem with the new Item. Returns the new Item.
func (l *List) Push(value interface{}) *Item {
	i := NewItem(value)
	l.PushItem(i)
	return i
}
