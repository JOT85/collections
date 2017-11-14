package list

import "errors"

//ErrNoItems is thrown when a method expects a List to contain Items and it doesn't.
var ErrNoItems = errors.New("List contains no Items")

//ErrInList is thrown when an Item is required to not be in a List, but is.
var ErrInList = errors.New("Item not in list")

//ErrNotInList is thrown when an Item is required to be in a List, but isn't.
var ErrNotInList = errors.New("Item not in list")

//ErrIndexOutOfRange is thrown when something is attempted to be carries out on an index that doesn't exist.
var ErrIndexOutOfRange = errors.New("index out of range")
