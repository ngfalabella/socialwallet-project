package aplication

import "social-wallet/domain"

func Transfer( fromWallet *domain.Wallet ,  toWallet *domain.Wallet  , amount int) (domain.Transaction , error ){
	if fromWallet.WalletOwnerID == toWallet.WalletOwnerID {
		return domain.Transaction{} , ErrSameWalletOwner
	}
	if fromWallet.WalletCurrency != toWallet.WalletCurrency {
		return domain.Transaction{} , ErrCurrencyMismatch
	}
	transactionResult , err := domain.NewTransactionConstructor(fromWallet.WalletOwnerID ,toWallet.WalletOwnerID, amount, fromWallet.WalletCurrency )
  		
	if err != nil {
		return domain.Transaction{} , err
	}
	
	err = fromWallet.Withdraw(amount) 
	if err != nil {
		return domain.Transaction{} , err
	}

	err = toWallet.Deposit(amount)

	if err != nil {
		return domain.Transaction{} , err
	}
	
	return transactionResult , nil
}