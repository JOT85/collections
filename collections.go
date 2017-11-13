//Package collections contains implementations of many collection data types for Go
//Made in a couple of computer science lessons, and a bit for fun too.
package collections

import "github.com/jot85/collections/list"
import "github.com/jot85/collections/queue"
import "github.com/jot85/collections/priorityQueue"
import "github.com/jot85/collections/stack"

//List is list.List from github.com/jot85/collections/list
type List list.List
//ListItem is list.Item from github.com/jot85/collections/list
type ListItem list.Item
//Queue is queue.Queue from github.com/jot85/collections/queue
type Queue queue.Queue
//PriorityQueue is priorityQueue.PriorityQueue from github.com/jot85/collections/priorityQueue
type PriorityQueue priorityQueue.PriorityQueue
//Stack is stack.Stack from github.com/jot85/collections/stack
type Stack stack.Stack
