package domain

type Transaction struct {
	FromUserID int
	ToUserID   int
	Amount     int
}


func NewTransaction(fromUserID int , toUserID int , amount int) Transaction {
	return Transaction{
		FromUserID: fromUserID,
		ToUserID: toUserID,
		Amount: amount,
	}
}