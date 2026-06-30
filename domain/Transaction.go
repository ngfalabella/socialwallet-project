package domain

type Transaction struct {
	UserFromID          int
	UserToID            int
	TransactionAmount   int
	TransactionCurrency Currency
}

func NewTransaction(from int, to int, amount int, currency Currency) (Transaction, error) {
	if from == to {
		return Transaction{}, InvalidUsersTransaction
	}
	if amount <= 0 {
		return Transaction{} ,  InvalidAmountTransaction
	}
	return Transaction{
		UserFromID: from,
		UserToID: to,
		TransactionAmount: amount,
		TransactionCurrency: currency,
	} , nil
}