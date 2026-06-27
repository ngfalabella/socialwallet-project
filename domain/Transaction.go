package domain

type Transaction struct {
	FromUserID  int
	ToUserID int
	Amout int
}

func NewTransaction(fromUserid int , toUserID int , amout int ) ( Transaction , error) {
	if amout <= 0 {
		return Transaction{} , CeroValueError
	} 
	if fromUserid == toUserID {
		return Transaction{} , IDEqualsError
	}
	return Transaction{
		FromUserID: fromUserid,
		ToUserID: toUserID,
		Amout: amout,
	},nil
}

func Transfer ( fromWallet *Wallet , toWallet *Wallet , amount int ) (Transaction , error) {
	if amount <= 0 {
		return Transaction{} , CeroValueError
	}
	if fromWallet.OwnerID  == toWallet.OwnerID {
		return Transaction{} , IDEqualsError
	}
	if fromWallet.Currency != toWallet.Currency {
		return Transaction{} , DiferentCurrencyErrors
	}
	if fromWallet.Balance < amount {
		return Transaction{} , AmountInsuficient
	}

	fromWallet.
}