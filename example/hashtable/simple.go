package main

import "github.com/jot85/collections"
import "github.com/jot85/collections/hashtable"
import "fmt"
import "unsafe"

func doHash(data unsafe.Pointer, attempt, last uint64) uint64 {
	if attempt > 0 {
		return (last + 1) % 10
	}
	fmt.Println(uintptr(data))
	fmt.Println(*(*uintptr)(data))
	fmt.Println(*(*[64]byte)(data))
	fmt.Println(*(*[5]byte)(
		unsafe.Pointer(
			*(*uintptr)(data),
		),
	))
	//fmt.Println(*(*[]byte)(data))
	n := *(*uint64)(data)
	return n % 10
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
	t := hashtable.NewHashTable(containerInterface, doHash, nil, true)
	printContainer(container)
	/*addit := func(v uint64) {
		t.AddPointer(unsafe.Pointer(&v))
	}*/
	//addit(5)
	//addit(7)
	//addit(18)
	printContainer(container)
	//addit(107)
	printContainer(container)
	temp := "Hi?"
	temp2 := "Yo!!!"
	t.AddPointer(unsafe.Pointer(&temp))
	t.AddPointer(unsafe.Pointer(&temp2))
	printContainer(container)
}
