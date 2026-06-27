package domain


type Wallet struct {
	OwnerID int 
	Balance int
	Currency string
}

func NewWallet(ownerID int , currency string ) Wallet {
	return Wallet {
		OwnerID: ownerID,
		Balance: 0,
		Currency: currency,
	}
}

func ( wallet *Wallet) Deposit( amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	wallet.Balance += amount
	return nil
}

func ( wallet  *Wallet) Withdraw( amount int) error {

	if amount <= 0{
		return ErrInvalidAmount
	}

	if amount > wallet.Balance {
		return ErrInsufficientBalance
	}
		wallet.Balance -= amount
		return nil
}