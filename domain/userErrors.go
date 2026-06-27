package domain

import "errors" 

var (
	ErrInvalidEmail = errors.New("El email es invalido, no puede estar vacio")
)