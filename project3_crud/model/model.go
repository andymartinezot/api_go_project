package model

import "errors"

var (
	//ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("The person can't be null")
	//ErrIDPersonDoesnotExist la persona no puede ser nula
	ErrIDPersonDoesnotExists = errors.New("The person doesn't exist")
)
