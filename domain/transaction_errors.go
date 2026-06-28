package domain

import "errors"

var (
	ErrSameUserTransaction = errors.New("No puede ser el mismo usuario")
	ErrInvalidTransactionAmount = errors.New("Monto de operacion invalida")
)