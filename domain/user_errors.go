package domain

import "errors"

var (
	InvalidNameUser = errors.New("El nombre de usuario no puede estar vacio")
	InvalidEmailUser = errors.New("El email no puede estar vacio")
)