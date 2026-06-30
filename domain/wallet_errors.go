package domain

import "errors"

var (
	InvalidCurrencyWallet = errors.New("Moneda Invalida - La moneda que seleccionaste no es compatible con la plataforma")
	InvalidMountWallet = errors.New("No podes usar montos negativos o iguales a 0")
	InvalidBalanceFinal = errors.New("El saldo es insuficiente")
)