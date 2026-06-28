package aplication

import "errors"

var (
	ErrSameWalletOwner = errors.New("No se puede usar la misma billetera")
	ErrCurrencyMismatch = errors.New("No se pueden hacer transacciones de distintas monedas")
)