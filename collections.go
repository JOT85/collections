//Package collections contains implementations of many collection data types for Go
//Made in a couple of computer science lessons, and a bit for fun too.
package collections

import "collections/list"
import "collections/queue"
import "collections/priorityQueue"
import "collections/stack"

//List is list.List from collections/list
type List list.List
//ListItem is list.Item from collections/list
type ListItem list.Item
//Queue is queue.Queue from collections/queue
type Queue queue.Queue
//PriorityQueue is priorityQueue.PriorityQueue from collections/priorityQueue
type PriorityQueue priorityQueue.PriorityQueue
//Stack is stack.Stack from collections/stack
type Stack stack.Stack
