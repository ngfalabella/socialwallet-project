package domain

import "errors"

var (
	EmptyFieldTransactionError = errors.New("Ningun campo puede estar vacio")
	AmountInsuficient = errors.New("No tienes suficiente dinero para hacer la transaccion")
	CeroValueError = errors.New("No puede ser 0 el valor de la transaccion")
	IDEqualsError =errors.New("No pueden ser ambos ID`s iguales")
	DiferentCurrencyErrors = errors.New("No pueden tener distinta moneda")
)