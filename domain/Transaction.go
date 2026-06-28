package domain

type Transaction struct {
	TransactionFromUserID int
	TransactionToUserID int
	TransactionAmount int
	TransactionCurrency Currency
}

func NewTransactionConstructor ( fromUserID int , toUserID int, amount int , currency Currency) (Transaction,error) {
	if fromUserID == toUserID {
		return Transaction{} , ErrSameUserTransaction
	}
	if amount <= 0 {
		return Transaction{} , ErrInvalidTransactionAmount
	}

	return Transaction{
			TransactionFromUserID : fromUserID ,
			TransactionToUserID :toUserID,
			TransactionAmount :amount,
			TransactionCurrency : currency,
	} , nil
}