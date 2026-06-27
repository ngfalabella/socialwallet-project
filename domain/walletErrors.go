package domain

import "errors"

var (
	ErrInvalidAmount = errors.New("El monto debe ser mayor a cero")
	ErrInsufficientBalance =  errors.New("Saldo insuficiente")
)