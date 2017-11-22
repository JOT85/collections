package list

import "unsafe"

//Item represents an item in a list.
type Item struct {
	value  unsafe.Pointer
	before *Item
	after  *Item
}

//NewItemFromPointer returns a pointer to a new Item of which the value is the value of the given unsafe.Pointer.
func NewItemFromPointer(value unsafe.Pointer) *Item {
	return &Item{value, nil, nil}
}

//NewItem returns a pointer to a new Item of which the value is the value given.
func NewItem(value interface{}) *Item {
	return &Item{unsafe.Pointer(&value), nil, nil}
}

//Copy returns a copy of the Item (not a part of any List though).
func (i *Item) Copy() Item {
	return Item{
		i.value,
		nil, nil,
	}
}

//CopyTo clones the source Item to the given Item.
func (i *Item) CopyTo(item *Item) {
	item.value = i.value
}

//ValuePointer returns the unsafe.Pointer to the value of this item.
func (i *Item) ValuePointer() unsafe.Pointer {
	return i.value
}

//Value returns the value of this item as an interface{}.
func (i *Item) Value() interface{} {
	return *(*interface{})(i.value)
}

//PointerValueOf (item) is the same as item.ValuePointer()
func (l *List) PointerValueOf(item *Item) unsafe.Pointer {
	return item.value
}

//ValueOf (item) is the same as item.Value()
func (l *List) ValueOf(item *Item) interface{} {
	return item.Value()
}

//SetPointer sets the value of the ListItem to the value in the given unsafe.Pointer.
func (l *List) SetPointer(i *Item, value unsafe.Pointer) {
	l.PanicIfUnsafeNotAllowed()
	i.value = value
}

//Set wraps SetPointer by calling SetPointer with the pointer to the given value.
func (l *List) Set(i *Item, value interface{}) {
	i.value = l.valueToPointer(value)
}
