package hashtable

import "github.com/jot85/collections"
import "github.com/jot85/collections/list"
import "unsafe"

type HashFunction func (unsafe.Pointer, uint64, uint64) uint64

type HashTable struct {
	container collections.IndexableSetablePointers
	cmap containermap
	HashFunction HashFunction
}

type item struct {
	data unsafe.Pointer
	stored bool
}

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
	l.list.GetIndex(i).SetPointer(unsafe.Pointer(&v))
}

func NewHashTable(container collections.IndexableSetablePointers, hashFunc HashFunction) HashTable {
	var m containermap
	l := container.Length()
	if l > uint64((^uint(0)) >> 1) {
		m = containermap(&listmap{list.NewList(l)})
	} else {
		temp := slicemap(make([]bool, l))
		m = containermap(&temp)
	}
	return HashTable{
		container,
		m,
		hashFunc,
	}
}

func (table *HashTable) AddPointer(value unsafe.Pointer) {
	var pos uint64
	var attempt uint64
	for {
		pos = table.HashFunction(value, attempt, pos)
		if !table.cmap.Get(pos) {
			table.cmap.Set(pos, true)
			table.container.SetIndex(pos, value)
			break
		}
		attempt++
	}
}
