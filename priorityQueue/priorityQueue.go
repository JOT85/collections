package priorityQueue

import (
	"unsafe"
)

//itemContainer contains an item, it's priority, and a link to another itemContainer
type itemContainer struct {
	item     unsafe.Pointer
	priority int
	behind   *itemContainer
}

//MaxPriority is the highest possible priority value
const MaxPriority = int(^uint(0) >> 1)

//PriorityQueue represents a dynamic collection of items that all have a priority
type PriorityQueue struct {
	front *itemContainer
	back  *itemContainer
}

//EnqueuePointer adds an unsafe.Pointer in front of the first item that has a lower priority than the given priority.
//If priority < 0, the item is added to the back of the list. Of if priority == MaxPriority, it is added to the front of the list.
func (q *PriorityQueue) EnqueuePointer(item unsafe.Pointer, priority int) {
	//Create item to add
	i := itemContainer{
		item:     item,
		priority: priority,
		behind:   nil,
	}
	if priority < 0 {
		//Add to the back if priority < 0
		if q.back != nil {
			q.back.behind = &i
		} else {
			q.front = &i
		}
		q.back = &i
	} else if priority == MaxPriority {
		//Add to the front if we have MaxPriority
		i.behind = q.front
		if q.front == nil {
			q.back = &i
		}
		q.front = &i
	} else {
		//Go though all the items until we have a greater priority than the current item
		c := q.front
		var p *itemContainer
		for c != nil {
			if priority > c.priority {
				//Insert before c
				i.behind = c
				//And after p (or at the front if p is nik)
				if p != nil {
					p.behind = &i
				} else {
					q.front = &i
				}
				break
			}
			p = c
			c = c.behind
		}
		//Add the the back if we have got to the end of the list and haven't added it
		if c == nil {
			if p != nil {
				p.behind = &i
			} else {
				q.front = &i
				q.back = &i
			}
		}
	}
}

//Enqueue wraps EnqueuePointer by taking any value as an interface and calling EnqueuePointer with a pointer to the value
func (q *PriorityQueue) Enqueue(item interface{}, priority int) {
	q.EnqueuePointer(unsafe.Pointer(&item), priority)
}

//PeekPointer returns a unsafe.Pointer to the item at the front of the queue, and it's priority
func (q *PriorityQueue) PeekPointer() (unsafe.Pointer, int) {
	return q.front.item, q.front.priority
}

//Peek wraps PeekPointer by returning the interface{} value of the returned pointer
func (q *PriorityQueue) Peek() (interface{}, int) {
	val, p := q.PeekPointer()
	return *(*interface{})(val), p
}

//DequeuePointer removes the item at the front of the queue and returns an unsafe.Pointer to the value, and the items priority
func (q *PriorityQueue) DequeuePointer() (unsafe.Pointer, int) {
	//Get the item
	i := q.front
	//Move the front one back
	q.front = i.behind
	//If the front is nil, then the back need to also be nil
	if q.front == nil {
		q.back = nil
	}
	return i.item, i.priority
}

//Dequeue wraps DequeuePointer by returning the interface{} value of the returned pointer
func (q *PriorityQueue) Dequeue() (interface{}, int) {
	val, p := q.DequeuePointer()
	return *(*interface{})(val), p
}

//IsEmpty returns true if there are no items in queue, false if otherwise
func (q *PriorityQueue) IsEmpty() bool {
	return q.front == nil
}

//Length returns the amount of items in the queue
func (q *PriorityQueue) Length() (l int) {
	i := q.front
	//While there is an item, inc l and go to the next item
	for i != nil {
		l++
		i = i.behind
	}
	return
}

//Count counts the number of items with the specified priority, if includeHigher is true, it also counts items with higher priorities
func (q *PriorityQueue) Count(priority int, includeHigher bool) (count int) {
	i := q.front
	//While there is an item
	for i != nil {
		//Break if the priority is lower than the search, otherwise inc count if nessisery
		if i.priority < priority {
			break
		} else if includeHigher {
			count++
		} else if i.priority == priority {
			count++
		}
		i = i.behind
	}
	return
}
