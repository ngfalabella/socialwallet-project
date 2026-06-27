package main

import (
	"fmt"
	"social-wallet/domain"
)

func main() {
	fmt.Println("*** APP FINTECH WALLET ***")
	fmt.Printf("\n")

	fmt.Println("*** Se crean los Usuarios ***")
	userGabriel := domain.User{ID: 1 ,Name: "Gabriel",Email: "gabriel@email.com"}
	userAna := domain.User{ID: 2 ,Name: "Ana",Email: "ana@email.com"}
	fmt.Println("- Gabriel : " , userGabriel)
	fmt.Println("- Ana : " , userAna)
	fmt.Printf("\n")

	fmt.Println("*** Se crean las Wallets ***")
	walletGabriel , err := domain.NewWallet(userGabriel.ID ,domain.CurrencyARS )
	if err != nil {
		fmt.Println("Error al crear wallet de Gabriel" , err)
	}
	fmt.Println("- Gabriel : " , walletGabriel)

	walletAna , err := domain.NewWallet(userAna.ID ,domain.CurrencyUSD )
	if err != nil {
		fmt.Println("Error al crear wallet de Ana" , err)
	}
	fmt.Println("- Ana : " , walletAna)

	

	
}