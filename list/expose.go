package list

import "unsafe"

//ExposedList is a struct that exposes all internal struct fields of a List
//You generally shouldn't use the exposed struct.
type ExposedList struct {
	Start  **Item
	End    **Item
	Length *uint64
	List   *List
}

//Expose creates an ExposedList from the List.
func (l *List) Expose() ExposedList {
	return ExposedList{
		Start:  &l.start,
		End:    &l.end,
		Length: &l.length,
		List:   l,
	}
}

//ExposedItem is a struct that exposes all internal struct fields of an Item
//You generally shouldn't use the exposed struct.
type ExposedItem struct {
	Before **Item
	After  **Item
	Value  *unsafe.Pointer
	Item   *Item
}

//Expose creates an ExposedItem from the Item.
func (i *Item) Expose() ExposedItem {
	return ExposedItem{
		Before: &i.before,
		After:  &i.after,
		Value:  &i.value,
		Item:   i,
	}
}
