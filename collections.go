//Package collections contains implementations of many collection data types for Go
//Made in a couple of computer science lessons, and a bit for fun too.
package collections

import "./list"
import "./queue"
import "./priorityQueue"
import "./stack"

//List is list.List from ./list
type List list.List
//ListItem is list.Item from ./list
type ListItem list.Item
//Queue is queue.Queue from ./queue
type Queue queue.Queue
//PriorityQueue is priorityQueue.PriorityQueue from ./priorityQueue
type PriorityQueue priorityQueue.PriorityQueue
//Stack is stack.Stack from ./stack
type Stack stack.Stack
