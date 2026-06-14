package main

import "fmt"

func main() {

	var senderName string = "Gabriel"
	var receiverName string = "Ana"
	amountInCents := 2500
	senderBalanceInCents := 10000
	receiverBalanceInCents := 2500
	operationID := "TRX-001"
	var isSuccessful bool
	isSuccessful = true

	senderBalanceAfterInCents := senderBalanceInCents - amountInCents
	receiverBalanceAfterInCents := receiverBalanceInCents + amountInCents

	var summary string = fmt.Sprintf("Resumen: %s transfirio %d centavos a %s", senderName, amountInCents, receiverName)

	fmt.Println("=== SocialWallet Transfer ===")
	fmt.Println("Operación:", operationID)
	fmt.Println("Origen:", senderName)
	fmt.Println("Destino:", receiverName)
	fmt.Println("Monto:", amountInCents, "centavos")
	fmt.Printf("Saldo origen antes: %d centavos\n", senderBalanceInCents)
	fmt.Println("Saldo destino antes:", receiverBalanceInCents, "centavos")
	fmt.Println("Transferencia exitosa:", isSuccessful)
	fmt.Println(summary)
	fmt.Printf("Saldo origen después:%d centavos \n", senderBalanceAfterInCents)
	fmt.Printf("Saldo destino después:%d centavos \n", receiverBalanceAfterInCents)
}
