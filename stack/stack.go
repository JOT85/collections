package stack

import "unsafe"

//itemContainer contains a pointer to an item and a pointer to the itemContainer under it.
type itemContainer struct {
	item  unsafe.Pointer
	under *itemContainer
}

//Stack stores a synamic collection if items in a FILO way
type Stack struct {
	top *itemContainer
}

//PushPointer adds the given pointer to the top of the stack
func (s *Stack) PushPointer(item unsafe.Pointer) {
	//Create the item, linking to the current top item
	i := itemContainer{
		item:  item,
		under: s.top,
	}
	//And set it as the new top
	s.top = &i
}

//Push wraps PushPointer by taking any value as an interface and calling PushPointer with a pointer to the value
func (s *Stack) Push(item interface{}) {
	s.PushPointer(unsafe.Pointer(&item))
}

//PeekPointer returns an unsafe.Pointer to the item on the top of the stack
func (s *Stack) PeekPointer() unsafe.Pointer {
	return s.top.item
}

//Peek wraps PeekPointer by returning the interface{} value of the returned pointer
func (s *Stack) Peek() interface{} {
	return *(*interface{})(s.PeekPointer())
}

//PopPointer removes the value on the top of the stack, and returns an unsafe.Pointer of the value
func (s *Stack) PopPointer() unsafe.Pointer {
	//Get the top value
	i := s.top
	//Set the top to the one under it
	s.top = i.under
	return i.item
}

//Pop wraps PopPointer by returning the interface{} value of the returned pointer
func (s *Stack) Pop() interface{} {
	return *(*interface{})(s.PopPointer())
}

//IsEmpty returns true if there are no items in stack, false if otherwise
func (s Stack) IsEmpty() bool {
	return s.top == nil
}

//Length returns the amount of items in the stack
func (s *Stack) Length() (l int) {
	i := s.top
	//While there is an item, inc l and go to the next item
	for i != nil {
		l++
		i = i.under
	}
	return
}
