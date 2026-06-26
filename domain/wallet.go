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

func ( wallet *Wallet) Deposit( amount int) bool {
	if amount <= 0 {
		return false
	}
	wallet.Balance += amount
	return true
}
func ( wallet  *Wallet) Withdraw( amount int) bool {

	if amount <= 0{
		return false
	}

	if amount > wallet.Balance {
		return false
	}
		wallet.Balance -= amount
		return true
}