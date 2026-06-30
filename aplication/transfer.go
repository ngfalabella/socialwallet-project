package aplication

import (
	"fmt"
	"social-wallet/domain"
)

func TransferOperation(walletOwner *domain.Wallet, walletTo *domain.Wallet, amount int) (domain.Transaction, error) {
	if walletOwner.WalletID == walletTo.WalletID {
		return domain.Transaction{}, InvalidTransferOwner
	}
	if walletOwner.WalletCurrency != walletTo.WalletCurrency {
		return domain.Transaction{}, InvalidCurrencyTransfer
	}
	transactionFinal , err := domain.NewTransaction(walletOwner.WalletID, walletTo.WalletID, amount, walletOwner.WalletCurrency)
	if err != nil {
		return domain.Transaction{}, err
	}

	err = walletOwner.Withdraw(amount)

	if err != nil {
		return domain.Transaction{} , fmt.Errorf(
			"no se pudo retirar el dinero de la wallet de origen : %w" , err,
		)
	}

	err = walletTo.Deposit(amount)

	if err != nil {
		return domain.Transaction{} ,  fmt.Errorf(
			"no se pudo depositar el dinero en la wallet destino : %w", err ,
		)
	}

	return transactionFinal , nil
}
