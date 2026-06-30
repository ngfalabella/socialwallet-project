package domain

import "errors"

var (
	InvalidUsersTransaction = errors.New("No pueden hacerse transacciones al mismo ID usuario")
	InvalidAmountTransaction = errors.New("No se pueden registrar Transacciones en monto negativo")
)