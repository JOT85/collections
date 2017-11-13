//Package list implements a linked list, conatining pointers to values.
package list

import (
	"errors"
	"unsafe"
)

//ErrNoItems is thrown when a method expects a List to contain Items and it doesn't.
var ErrNoItems = errors.New("List contains no Items")

//ErrInList is thrown when an Item is required to not be in a List, but is.
var ErrInList = errors.New("Item not in list")

//ErrNotInList is thrown when an Item is required to be in a List, but isn't.
var ErrNotInList = errors.New("Item not in list")

//ErrIndexOutOfRange is thrown when something is attempted to be carries out on an index that doesn't exist.
var ErrIndexOutOfRange = errors.New("index out of range")

//List type represents a dynamic queue that stores unsafe.Pointers
type List struct {
	start  *Item
	end    *Item
	length uint64
}

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
	if item.before != nil && item.after != nil {
		panic(errors.New("Item to insert relitive to must be part of a List"))
	}
	l.insertBefore(item, insert)
}

//InsertItemAfter inserts the given item (must not be part of a list) after the given Item in this List.
func (l *List) InsertItemAfter(item, insert *Item) {
	if insert.before != nil && insert.after != nil {
		panic(errors.New("Item to insert must not be part of a List"))
	}
	if item.before != nil && item.after != nil {
		panic(errors.New("Item to insert relitive to must be part of a List"))
	}
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

//ShiftItem inserts the Item at the start of the list.
func (l *List) ShiftItem(item *Item) {
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

//ShiftPointer creates a new Item containing the pointer and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) ShiftPointer(value unsafe.Pointer) *Item {
	item := &Item{
		value,
		nil, nil,
	}
	l.ShiftItem(item)
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

//Shift creates a new Item containing the given value and calls ShiftItem with the new Item. Returns the new Item.
func (l *List) Shift(value interface{}) *Item {
	return l.ShiftPointer(unsafe.Pointer(&value))
}

//Push creates a new Item containing the given value and calls PushItem with the new Item. Returns the new Item.
func (l *List) Push(value interface{}) *Item {
	return l.PushPointer(unsafe.Pointer(&value))
}

//Pop removes and returns the last Item in the List
func (l *List) Pop() *Item {
	i := l.GetIndex(l.length-1)
	l.remove(i)
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
