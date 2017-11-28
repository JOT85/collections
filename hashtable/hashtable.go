package hashtable

import "github.com/jot85/collections"
import "github.com/jot85/collections/list"
import "unsafe"
import "errors"

var ErrFull = errors.New("HashTable container full")

//HashFunction represents a function that can be used as a hash function for a hash table.
//The first argument is a pointer to the data it should hash.
//The second is the current hash attempt - this gets incremented on a collision.
//And if the second argument is greater that 0 (so a collision has occured), the final argument is the last hashed value.
//It should return the hash.
type HashFunction func(unsafe.Pointer, uint64, uint64) uint64

//HashTable represents a hash table.
type HashTable struct {
	container    collections.IndexableSetablePointers
	cmap         containermap
	HashFunction HashFunction
	allowUnsafe  bool
	setter       collections.SetFunction
	getter       collections.GetFunction
}

type item struct {
	data   unsafe.Pointer
	stored bool
}

//NewHashTable creates a new HashTable using the given collections.IndexableSetablePointers to store the data in, and the given HashFunction to generate the hashes.
func NewHashTable(container collections.IndexableSetablePointers, hashFunc HashFunction, setter collections.SetFunction, getter collections.GetFunction, allowUnsafe bool) HashTable {
	var m containermap
	l := container.Length()
	if l > uint64((^uint(0))>>1) {
		m = containermap(&listmap{list.NewList(l, nil, true)})
	} else {
		temp := slicemap(make([]bool, l))
		m = containermap(&temp)
	}
	return HashTable{
		container,
		m,
		hashFunc,
		allowUnsafe,
		setter,
		getter,
	}
}

func (table *HashTable) valueToPointer(value interface{}) unsafe.Pointer {
	if table.setter == nil {
		return unsafe.Pointer(&value)
	}
	return table.setter(value)
}

//addPointer adds the value pointed to by the given unsafe.Pointer to the HashTable
func (table *HashTable) addPointer(value unsafe.Pointer) {
	var pos uint64
	var attempt uint64
	pos = table.HashFunction(value, attempt, pos)
	origPos := pos
	for {
		if !table.cmap.Get(pos) {
			table.cmap.Set(pos, true)
			table.container.SetIndex(pos, value)
			break
		}
		attempt++
		pos = table.HashFunction(value, attempt, pos)
		if pos == origPos {
			panic(ErrFull)
		}
	}
}

//AddPointer adds the value pointed to by the given unsafe.Pointer to the HashTable
func (table *HashTable) AddPointer(value unsafe.Pointer) {
	table.addPointer(value)
}

//Add adds the value to the HashTable
func (table *HashTable) Add(value interface{}) {
	table.addPointer(table.valueToPointer(value))
}

func (table *HashTable) Search(value interface{}) bool {
	var pos uint64
	var attempt uint64
	pVal := table.valueToPointer(value)
	pos = table.HashFunction(pVal, attempt, pos)
	origPos := pos
	for {
		if !table.cmap.Get(pos) {
			return false
		}
		val := table.getter(table.container.GetIndex(pos))
		if value == val {
			return true
		}
		attempt++
		pos = table.HashFunction(pVal, attempt, pos)
		if pos == origPos {
			return false
		}
	}
}

func (table *HashTable) SearchPointer(value unsafe.Pointer) bool {
	return table.Search(table.setter(value))
}
