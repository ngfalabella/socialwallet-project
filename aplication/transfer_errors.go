package aplication

import "errors"

var (
	InvalidTransferOwner = errors.New("Transferencias invalidas entre wallets del mismo propietario")
	InvalidCurrencyTransfer = errors.New("No se pueden registrar tranferencias de distintas monedas")
	InvalidZeroValue = errors.New("El valor de la transferencia no puede ser igual a 0")
)