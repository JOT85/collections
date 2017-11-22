package list

import "unsafe"

//insertAssumingEmpty inserts the Item at index 0 assuming there are no Items in the List.
func (l *List) insertAssumingEmpty(insert *Item) {
	//The length is now 1, and this Item is the start and the end
	l.length = 1
	l.start = insert
	l.end = insert
	l.lastIndexedIndex = 0
	l.lastIndexedItem = insert
}

//insertBefore inserts the second Item before ther first Item in the List.
func (l *List) insertBefore(item, insert *Item) {
	l.length++
	l.lastIndexedIndex++
	if l.isFirstItem(item) {
		l.start = insert
	} else {
		item.before.after = insert
	}
	insert.before = item.before
	insert.after = item
	item.before = insert
}

//InsertItemBefore inserts the given item (must not be part of a list) before the given Item in this List.
func (l *List) InsertItemBefore(item, insert *Item) {
	l.PanicIfUnsafeNotAllowed()
	insert.PanicIfInList()
	l.PanicIfNotInList(item)
	l.insertBefore(item, insert)
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

//InsertItemAfter inserts the given item (must not be part of a list) after the given Item in this List.
func (l *List) InsertItemAfter(item, insert *Item) {
	l.PanicIfUnsafeNotAllowed()
	insert.PanicIfInList()
	l.PanicIfNotInList(item)
	l.insertAfter(item, insert)
}

//insertPointerBefore inserts a new item, holding the value given, before the Item. And returns a pointer to the new Item.
func (l *List) insertPointerBefore(item *Item, value unsafe.Pointer) *Item {
	ni := &Item{
		value,
		nil, nil,
	}
	l.insertBefore(item, ni)
	return ni
}

//InsertPointerBefore inserts a new item, holding the value given, before the Item. And returns a pointer to the new Item.
func (l *List) InsertPointerBefore(item *Item, value unsafe.Pointer) *Item {
	l.PanicIfUnsafeNotAllowed()
	return l.insertPointerBefore(item, value)
}

//insertPointerAfter inserts a new item, holding the value given, after the Item. And returns a pointer to the new Item.
func (l *List) insertPointerAfter(item *Item, value unsafe.Pointer) *Item {
	ni := &Item{
		value,
		nil, nil,
	}
	l.insertAfter(item, ni)
	return ni
}

//InsertPointerAfter inserts a new item, holding the value given, after the Item. And returns a pointer to the new Item.
func (l *List) InsertPointerAfter(item *Item, value unsafe.Pointer) *Item {
	l.PanicIfUnsafeNotAllowed()
	return l.insertPointerBefore(item, value)
}

//InsertBefore wraps InsertPointerBefore by calling InsertPointerBefore with a pointer to the given value.
func (l *List) InsertBefore(item *Item, value interface{}) *Item {
	return l.insertPointerBefore(item, l.valueToPointer(value))
}

//InsertAfter wraps InsertPointerAfter by calling InsertPointerAfter with a pointer to the given value.
func (l *List) InsertAfter(item *Item, value interface{}) *Item {
	return l.insertPointerAfter(item, l.valueToPointer(value))
}

//unshiftItem inserts the Item at the start of the list.
func (l *List) unshiftItem(item *Item) {
	if l.length == 0 {
		l.insertAssumingEmpty(item)
	} else {
		l.insertBefore(l.start, item)
	}
}

//UnshiftItem inserts the Item at the start of the list.
//Panics with ErrUnsafe if allowUnsafe is false
func (l *List) UnshiftItem(item *Item) {
	l.PanicIfUnsafeNotAllowed()
	l.unshiftItem(item)
}

//pushItem inserts the Item at the end of the list.
func (l *List) pushItem(item *Item) {
	if l.length == 0 {
		l.insertAssumingEmpty(item)
	} else {
		l.insertAfter(l.end, item)
	}
}

//PushItem inserts the Item at the end of the list.
//Panics with ErrUnsafe if allowUnsafe is false
func (l *List) PushItem(item *Item) {
	l.PanicIfUnsafeNotAllowed()
	l.pushItem(item)
}

//unshiftPointer creates a new Item containing the pointer and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) unshiftPointer(value unsafe.Pointer) *Item {
	item := &Item{
		value,
		nil, nil,
	}
	l.unshiftItem(item)
	return item
}

//UnshiftPointer creates a new Item containing the pointer and calls ShiftItem with the new Item. Returns the new Item.
//Panics with ErrUnsafe if allowUnsafe is false
func (l *List) UnshiftPointer(value unsafe.Pointer) *Item {
	l.PanicIfUnsafeNotAllowed()
	return l.unshiftPointer(value)
}

//pushPointer creates a new Item containing the pointer and calls PushItem with the new Item. Returns the new Item.
func (l *List) pushPointer(value unsafe.Pointer) *Item {
	item := &Item{
		value,
		nil, nil,
	}
	l.pushItem(item)
	return item
}

//PushPointer creates a new Item containing the pointer and calls PushItem with the new Item. Returns the new Item.
//Panics with ErrUnsafe if allowUnsafe is false
func (l *List) PushPointer(value unsafe.Pointer) *Item {
	l.PanicIfUnsafeNotAllowed()
	return l.pushPointer(value)
}

//Unshift creates a new Item containing the given value and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) Unshift(value interface{}) *Item {
	return l.unshiftPointer(l.valueToPointer(value))
}

//Push creates a new Item containing the given value and calls PushItem with the new Item. Returns the new Item.
func (l *List) Push(value interface{}) *Item {
	return l.pushPointer(l.valueToPointer(value))
}
