package hashtable

import (
	"unsafe"

	"github.com/jot85/collections/list"
)

type containermap interface {
	Get(uint64) bool
	Set(uint64, bool)
}

type slicemap []bool

func (m *slicemap) Get(i uint64) bool {
	return (*m)[int(i)]
}

func (m *slicemap) Set(i uint64, v bool) {
	(*m)[int(i)] = v
}

type listmap struct {
	list *list.List
}

func (l *listmap) Get(i uint64) bool {
	return *(*bool)(l.list.GetIndex(i).ValuePointer())
}

func (l *listmap) Set(i uint64, v bool) {
	l.list.SetPointer(l.list.GetIndex(i), unsafe.Pointer(&v))
}
