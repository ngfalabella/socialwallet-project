package main

import (
	"fmt"
	"social-wallet/domain"
	)

func main() {
	userGabriel := domain.NewUser(1,"Gabriel","gabriel@gmail.com")
	userAna := domain.NewUser(2,"Ana","ana@gmail.com")

	walletGabriel := domain.NewWallet(userGabriel.ID ,"ARS")
	walletAna := domain.NewWallet(userAna.ID ,"ARS")


	transactionGabriel := domain.NewTransaction(userGabriel.ID, userAna.ID,2500)

	fmt.Println("Usuarios : ")
	fmt.Println("GABRIEL : " ,userGabriel)
	fmt.Println("ANA : " ,userAna)
	
	fmt.Printf("\n")
	
	fmt.Println("Wallets : ")
	fmt.Println("GABRIEL : " ,walletGabriel)
	fmt.Println("ANA : " ,walletAna)

	fmt.Printf("\n")
	fmt.Println("Transaction : ")
	fmt.Println("GABRIEL : " ,transactionGabriel)

	fmt.Println("Antes del deposito " , walletGabriel.Balance )
	walletGabriel.Deposit(10000) 
	fmt.Println("Despues del deposito " , walletGabriel.Balance )

}