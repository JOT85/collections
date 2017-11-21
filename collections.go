//Package collections contains implementations of many collection data types for Go
//Made in a couple of computer science lessons, and a bit for fun too.
package collections

import "github.com/jot85/collections/list"
import "unsafe"

type IndexableSetablePointers interface {
	GetIndex(uint64) unsafe.Pointer
	SetIndex(uint64, unsafe.Pointer)
	Length() uint64
}

type ipList struct {
	list *list.List
}

func (list *ipList) GetIndex(index uint64) unsafe.Pointer {
	return list.list.GetIndex(index).ValuePointer()
}

func (list *ipList) SetIndex(index uint64, value unsafe.Pointer) {
	list.list.GetIndex(index).SetPointer(value)
}

func (list *ipList) Length() uint64 {
	return list.list.Length()
}

type ipSlice struct {
	slice* []unsafe.Pointer
}

func (slice *ipSlice) GetIndex(index uint64) unsafe.Pointer {
	return slice.slice[int(index)]
}

func (slice *ipSlice) SetIndex(index uint64, value unsafe.Pointer) {
	slice.slice[int(index)] = value
}

func (slice *ipSlice) Length() uint64 {
	return uint64(len(slice.slice))
}

func NewIndexablePointers(storage *interface{}) IndexableSetablePointers {
	switch t := (storage).(type) {
		case list.List:
			return ipList{t}
		case []unsafe.Pointer:
			return ipSlice{t}
		default:
			panic("Not valid storage container (must be list.List or []unsafe.Pointer)")
	}
}
