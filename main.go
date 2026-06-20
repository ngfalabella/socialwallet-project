package main

import "fmt"

const operationWithdraw = "withdraw"
const operationDeposit = "deposit"
const operationTransfer = "transfer"

func processWalletOperation(balanceInCents int, amountInCents int, operation string) (int, bool) {

	if amountInCents <= 0 {
		return balanceInCents, false
	}

	if amountInCents > balanceInCents {
		return balanceInCents, false
	}

	switch operation {
	case operationWithdraw:
		return balanceInCents - amountInCents, true

	case operationDeposit:
		return balanceInCents + amountInCents, true

	case operationTransfer:
		return balanceInCents - amountInCents, true

	default:
		return balanceInCents, false

	}

}

func main() {

	userBalanceTest := 10000
	userAmounTest := 2500
	userOperationTest := operationDeposit

	result, ok := processWalletOperation(userBalanceTest, userAmounTest, userOperationTest)

	if !ok {
		fmt.Println("Ocurrio un error al procesar")
		fmt.Println("Saldo actual : ", result)
		return
	}

	fmt.Println("Saldo actual : ", result)

}
