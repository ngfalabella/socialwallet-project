package domain

type Wallet struct {
	WalletOwnerID int
	WalletBalance int
	WalletCurrency Currency
}

func NewWalletConstructor (owenerID int , currency Currency )(Wallet,error) {
	if !currency.IsValid() {
		return Wallet{} , ErrInvalidCurrency
	}
	return Wallet{
			WalletOwnerID : owenerID ,
			WalletBalance : 0 ,
			WalletCurrency : currency ,
	} , nil
}

func (wallet *Wallet) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	wallet.WalletBalance += amount
	return nil
}

func (wallet *Wallet) Withdraw(amount int) error {
	if amount == 0 {
		return ErrInvalidAmount
	}
	if amount > wallet.WalletBalance {
		return ErrInsufficientBalance
	}
	wallet.WalletBalance -= amount
	return nil
}