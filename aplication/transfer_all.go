package aplication

import "social-wallet/domain"

func TransferAllBalance(fromWallet *domain.Wallet , toWallet *domain.Wallet) (domain.Transaction , error) {
	if fromWallet.WalletBalance <= 0 {
		return domain.Transaction{} , InvalidZeroValue
	}
	
	allTransaction , err := TransferOperation(fromWallet , toWallet , fromWallet.WalletBalance)
	
	if err != nil {
		return domain.Transaction{} ,err
	}

	return allTransaction, nil

}