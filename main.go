package main

import "fmt"

func main() {

	var balanceInCents int = 15000
	var trasnferAmountInCents int = 2500
	var balanceAfterTransferInCents int = balanceInCents - trasnferAmountInCents


	var hasEnoughBalance bool

	hasEnoughBalance = balanceInCents >= trasnferAmountInCents

	fmt.Printf("Saldo suficiente:%t\n", hasEnoughBalance)
	fmt.Printf("Saldo :%d\n", balanceAfterTransferInCents)

}
