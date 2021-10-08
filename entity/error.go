package entity

import "errors"

var ErrNotFound = errors.New("not found")
var ErrDuplicatedItem = errors.New("item add before")
var ErrCartNotFound = errors.New("cart not found")
var ErrItemNotFound = errors.New("item not found")
