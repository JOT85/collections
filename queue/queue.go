package queue

import "unsafe"

//itemContainer contains a pointer to an item and a pointer to the itemContainer behind it.
type itemContainer struct {
	item   unsafe.Pointer
	behind *itemContainer
}

//Queue type represents a dynamic queue that stores unsafe.Pointers
type Queue struct {
	front *itemContainer
	back  *itemContainer
}

//EnqueuePointer adds the given unsafe.Pointer to the back of the queue
func (q *Queue) EnqueuePointer(item unsafe.Pointer) {
	//Item holder to add
	i := itemContainer{
		item:   item,
		behind: nil,
	}
	if q.back != nil {
		//Link the old back to this one
		q.back.behind = &i
	} else {
		//Otherwise, this is now the front
		q.front = &i
	}
	//And the back!
	q.back = &i
}

//Enqueue wraps EnqueuePointer by taking any value as an interface and calling EnqueuePointer with a pointer to the value
func (q *Queue) Enqueue(item interface{}) {
	q.EnqueuePointer(unsafe.Pointer(&item))
}

//PeekPointer returns a unsafe.Pointer to the item at the front of the queue.
func (q *Queue) PeekPointer() unsafe.Pointer {
	return q.front.item
}

//Peek wraps PeekPointer by returning the interface{} value of the returned pointer
func (q *Queue) Peek() interface{} {
	return *(*interface{})(q.PeekPointer())
}

//DequeuePointer removes the item at the front of the queue and returns an unsafe.Pointer to the value
func (q *Queue) DequeuePointer() unsafe.Pointer {
	//Removed item
	i := q.front.item
	//Set the front to the next item along
	q.front = q.front.behind
	//If the front is nil, the back should also be nil
	if q.front == nil {
		q.back = nil
	}
	return i
}

//Dequeue wraps DequeuePointer by returning the interface{} value of the returned pointer
func (q *Queue) Dequeue() interface{} {
	return *(*interface{})(q.DequeuePointer())
}

//IsEmpty returns true if there are no items in queue, false if otherwise
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

//Length returns the amount of items in the queue
func (q *Queue) Length() (l int) {
	i := q.front
	//While there is an item, inc l and go to the next item
	for i != nil {
		l++
		i = i.behind
	}
	return
}
