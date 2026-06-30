package domain

type Wallet struct {
	WalletID int
	WalletBalance int
	WalletCurrency Currency
}

func NewWallet( id int ,currency Currency) (Wallet , error ) {
	if !currency.IsValidCurrency() {
		return Wallet{},InvalidCurrencyWallet
	}
	return Wallet{
		WalletID: id,
		WalletBalance: 0,
		WalletCurrency: currency,
	},nil
}

func (wallet *Wallet) Deposit(amount int ) error {
	if amount <= 0 {
		return InvalidMountWallet
	}
	wallet.WalletBalance += amount
	return nil
}

func ( wallet *Wallet) Withdraw ( amount int) error {
	if amount  <= 0 {
		return InvalidMountWallet
	}
	if wallet.WalletBalance < amount {
		return InvalidBalanceFinal
	}
	wallet.WalletBalance -= amount
	return nil
}