package store

import "errors"

var (
	//ErrRecorcNotFound when record not found in a database
	ErrRecorcNotFound = errors.New("record not found")
)
