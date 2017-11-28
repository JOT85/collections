//Package collections contains implementations of many collection data types for Go
//Made in a couple of computer science lessons, and a bit for fun too.
package collections

import "unsafe"

//SetFunction represents a function that should be called with a value before it is inserted into a collection.
//It should return an unsafe.Pointer to be stored.
type SetFunction func(interface{}) unsafe.Pointer

type GetFunction func(unsafe.Pointer) interface{}

//IndexableSetablePointers is an interface something that allows indexing, has a length, and the values in the indexes can be retrieved or set.
//So slices, lists etc. can all be wrapped to fit this interface.
//One example of it's use is for containers in the "github.com/jot85/collections/hashtables" package.
type IndexableSetablePointers interface {
	GetIndex(uint64) unsafe.Pointer
	SetIndex(uint64, unsafe.Pointer)
	Length() uint64
}

type ipSlice struct {
	slice *[]unsafe.Pointer
}

//SliceToIndexableSetablePointers wraps a slice of unsafe.Pointers in a IndexableSetablePointers interface
func SliceToIndexableSetablePointers(slice *[]unsafe.Pointer) IndexableSetablePointers {
	return IndexableSetablePointers(&ipSlice{slice})
}

func (slice *ipSlice) GetIndex(index uint64) unsafe.Pointer {
	return (*slice.slice)[int(index)]
}

func (slice *ipSlice) SetIndex(index uint64, value unsafe.Pointer) {
	(*slice.slice)[int(index)] = value
}

func (slice *ipSlice) Length() uint64 {
	return uint64(len(*slice.slice))
}
