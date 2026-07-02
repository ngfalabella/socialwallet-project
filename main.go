package main

import (
	"fmt"
	"errors"
	"social-wallet/domain"
	"social-wallet/aplication"
)



func main() {
	userGabriel , err := domain.NewUser(1,"Gabriel","gabriel@gmail.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	UserAna , err := domain.NewUser(2 , "Ana" , "ana@gmail.com")
		
		if err != nil {
		fmt.Println(err)
		return
	}

  walletGabriel , err := domain.NewWallet(userGabriel.UserID , domain.CurrencyARS )

		if err != nil {
		fmt.Println(err)
		return
	}

	walletAna , err := domain.NewWallet(UserAna.UserID , domain.CurrencyARS )

		if err != nil {
		fmt.Println(err)
		return
	}

	err = walletGabriel.Deposit(10000)

		if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("***Saldos Gabriel y Ana ANTES de Transferir")
	fmt.Println("Gabriel : " , walletGabriel.WalletBalance)
	fmt.Println("Ana : " , walletAna.WalletBalance)

	transactionTest , err := aplication.TransferOperation(&walletGabriel,&walletAna,12000)
	
	if err != nil {
		if errors.Is(err ,domain.InvalidBalanceFinal){
			fmt.Println("La transferencia fue rechazada por saldo insuficiente")
			fmt.Println("Error completo:", err)
			return
		}
		fmt.Println("Ocurrió otro error:", err)
	  return
	}

	fmt.Println("***Saldos Gabriel y Ana DESPUES de Transferir")
	fmt.Println("Gabriel : " , walletGabriel.WalletBalance)
	fmt.Println("Ana : " , walletAna.WalletBalance)

	fmt.Printf("\n")

	fmt.Println("Transaccion Creada")
	fmt.Println(transactionTest)

}
