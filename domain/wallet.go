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

func ( wallet *Wallet) Deposit( amount int) {
	if amount <= 0 {
		return
	}
	wallet.Balance += amount
}