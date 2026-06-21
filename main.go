package main
import "fmt"

func main(){
	transactionAmounts := []int{1200, 2500, 3000, 1500, 4000}

	recentTransactions := transactionAmounts[2:]

	fmt.Println(recentTransactions)

	recentTransactions[0] = 9999

	fmt.Println(transactionAmounts)
	fmt.Println(recentTransactions)
}