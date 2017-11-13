package list

import (
	"unsafe"
)

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

//PanicIfInList panics with ErrInList if the Item is in a List
func (i *Item) PanicIfInList() {
	if i.before != nil || i.after != nil {
		panic(ErrInList)
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

//Copy returns a copy of the Item (not a part of any List though).
func (i *Item) Copy() Item {
	return Item{
		i.value,
		nil, nil,
	}
}
