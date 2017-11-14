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

//ValuePointer returns the unsafe.Pointer to the value of this item.
func (i *Item) ValuePointer() unsafe.Pointer {
	return i.value
}

//Value returns the value of this item as an interface{}.
func (i *Item) Value() interface{} {
	return *(*interface{})(i.value)
}

//SetPointer sets the value of the ListItem to the value in the given unsafe.Pointer.
func (i *Item) SetPointer(value unsafe.Pointer) {
	i.value = value
}

//Set wraps SetPointer by calling SetPointer with the pointer to the given value.
func (i *Item) Set(value interface{}) {
	i.value = unsafe.Pointer(&value)
}
