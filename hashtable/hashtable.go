package hashtable

import "github.com/jot85/collections"
import "github.com/jot85/collections/list"
import "unsafe"

type HashFunction func (unsafe.Pointer, uint64, uint64) uint64

type HashTable struct {
	container *collections.IndexableSetablePointers
	map map
	HashFunction HashFunction
}

type item struct {
	data unsafe.Pointer
	stored bool
}

type map interface{
	get func(uint64) bool
	set func(uint64, bool)
}

type slicemap []byte

func (m *slicemap) get(i uint64) bool {
	return m[int(i)]
}

func (m *slicemap) set(i uint64, v bool) {
	slicemap[int(i)] = v
}

type listmap list.List

func (l *listmap) get(i uint64) bool {
	return *(*bool)(l.GetIndex(i).ValuePonter())
}

func (l *listmap) set(i uint64, v bool) {
	l.GetIndex(i).SetPointer(unsafe.Pointer(&bool))
}

func NewHashTable(container collections.IndexablePointers, hashFunc HashFunction)) HashTable {
	var m map
	l := container.Length()
	if l > int((^uint(0)) >> 1) {
		m = map(listmap(list.NewList(len)))
	} else {
		m = map(slicemap(make([]bool, l)))
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
		if !table.map.get(pos) {
			table.map.set(pos, true)
			table.container.SetIndex(pos, value)
			break
		}
		attempt++
	}
}
