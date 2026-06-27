package domain

type Wallet struct  {
	OwnerID int
	Balance int
	Currency Currency
}


func NewWallet ( ownerID int , currencyMoney Currency) (Wallet,error) {
	if !currencyMoney.IsValid() {
		return Wallet{} , WalletTypeError
	}
	return Wallet{
		OwnerID: ownerID,
		Balance: 0,
		Currency: currencyMoney,
	} ,nil
}