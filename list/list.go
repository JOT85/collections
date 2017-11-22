//Package list implements a linked list, conatining pointers to values.
package list

import (
	"unsafe"

	"github.com/jot85/collections"
)

//List type represents a dynamic queue that stores unsafe.Pointers
type List struct {
	start            *Item
	end              *Item
	length           uint64
	setter           collections.SetFunction
	allowUnsafe      bool
	lastIndexedIndex uint64
	lastIndexedItem  *Item
}

//NewList returns a new list with length empty items.
func NewList(length uint64, setter collections.SetFunction, allowUnsafe bool) (l *List) {
	l.setter = setter
	for length > 0 {
		l.pushPointer(nil)
		length--
	}
	l.allowUnsafe = allowUnsafe
	return
}

//ToIndexableSetablePointers returns a collections.IndexableSetablePointers interface to the List
func (l *List) ToIndexableSetablePointers() collections.IndexableSetablePointers {
	return collections.IndexableSetablePointers(&ispList{l})
}

func (l *List) valueToPointer(value interface{}) unsafe.Pointer {
	if l.setter == nil {
		return unsafe.Pointer(&value)
	}
	return l.setter(value)
}

type ispList struct {
	list *List
}

func (list *ispList) GetIndex(index uint64) unsafe.Pointer {
	return list.list.GetIndex(index).ValuePointer()
}

func (list *ispList) SetIndex(index uint64, value unsafe.Pointer) {
	list.list.SetPointer(list.list.GetIndex(index), value)
}

func (list *ispList) Length() uint64 {
	return list.list.Length()
}
