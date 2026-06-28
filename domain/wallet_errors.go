package domain

import "errors"

var (
	ErrInvalidCurrency = errors.New("No ingresaste una moneda valida")
	ErrInvalidAmount = errors.New("El monto asignado es invalido")
	ErrInsufficientBalance = errors.New("El saldo es insuficiente")
)