package main

import "github.com/jot85/collections"
import "github.com/jot85/collections/hashtable"
import "fmt"
import "unsafe"

func doHash(data unsafe.Pointer, attempt, last uint64) uint64 {
	if attempt > 0 {
		return last + 1
	}
	n := *(*uint64)(data)
	return n%10
}

func printContainer(container []unsafe.Pointer) {
	for _, v := range container {
		fmt.Print(" ")
		if v == nil {
			fmt.Print("-")
		} else {
			fmt.Print(*(*uint64)(v))
		}
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	container := make([]unsafe.Pointer, 10)
	containerInterface := collections.NewIndexablePointers(&container)
	t := hashtable.NewHashTable(containerInterface, doHash)
	printContainer(container)
	addit := func (v uint64) {
		t.AddPointer(unsafe.Pointer(&v))
	}
	addit(5)
	addit(7)
	addit(18)
	printContainer(container)
	addit(107)
	printContainer(container)
}
