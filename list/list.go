//Package list implements a linked list, conatining pointers to values.
package list

//List type represents a dynamic queue that stores unsafe.Pointers
type List struct {
	start  *Item
	end    *Item
	length uint64
}

//NewList returns a new list with length empty items.
func NewList(length uint64) (l *List) {
	for length > 0 {
		l.PushPointer(nil)
		length--
	}
	return
}
