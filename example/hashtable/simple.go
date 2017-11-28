package main

import "github.com/jot85/collections"
import "github.com/jot85/collections/hashtable"
import "fmt"
import "unsafe"

func doHash(data unsafe.Pointer, attempt, last uint64) uint64 {
	if attempt > 0 {
		return (last + 1) % 10
	}
	//fmt.Println(uintptr(data))
	//fmt.Println(*(*uintptr)(data))
	//fmt.Println(*(*[64]byte)(data))
	/*fmt.Println(*(*[5]byte)(
		unsafe.Pointer(
			*(*uintptr)(data),
		),
	))*/
	//fmt.Println(*(*[]byte)(data))
	n := *(*uint64)(data)
	return n % 10
}

func yaySet(val interface{}) unsafe.Pointer {
	temp := val.(uint64)
	return unsafe.Pointer(&temp)
}

func yayGet(val unsafe.Pointer) interface{} {
	return *(*uint64)(val)
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
	fmt.Println([]byte("Hi?"))
	fmt.Println([]byte("Yo!!!"))
	container := make([]unsafe.Pointer, 10)
	containerInterface := collections.SliceToIndexableSetablePointers(&container)
	t := hashtable.NewHashTable(containerInterface, doHash, yaySet, yayGet, true)
	printContainer(container)
	t.Add(uint64(5))
	t.Add(uint64(7))
	t.Add(uint64(18))
	t.Add(uint64(107))
	t.Add(uint64(2148))
	printContainer(container)
	//temp := "Hi?"
	//temp2 := "Yo!!!"
	//t.AddPointer(unsafe.Pointer(&temp))
	//t.AddPointer(unsafe.Pointer(&temp2))
	printContainer(container)
	fmt.Println(t.Search(uint64(5)))
	fmt.Println(t.Search(uint64(7)))
	fmt.Println(t.Search(uint64(3)))
	fmt.Println(t.Search(uint64(8)))
	fmt.Println(t.Search(uint64(2148)))
}
