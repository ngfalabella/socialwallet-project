package main

import(
	 "fmt"
	 "social-wallet/domain"
	 "social-wallet/aplication"
)


func main() {
	UserGabriel , err  := domain.NewUserConstructor(1,"gabriel","gabriel@email.com")

	if err != nil {
		fmt.Println(err)		 
		return
	}

	UserAna , err  := domain.NewUserConstructor(2,"ana","ana@email.com")

	if err != nil {
		fmt.Println(err)		
		return 
	}

	WalletGabriel , err := domain.NewWalletConstructor(UserGabriel.UserID,domain.CurrencyARS )

		if err != nil {
		fmt.Println(err)
		return		 
	 }

	WalletAna , err := domain.NewWalletConstructor(UserAna.UserID,domain.CurrencyARS )

		if err != nil {
		fmt.Println(err)	
		return	 
	 }

	 err = WalletGabriel.Deposit(10000)

	 if err != nil {
		fmt.Println(err)		 
		return
	 }

	 fmt.Println("Saldo Actualizado de Gabriel : " , WalletGabriel.WalletBalance)
	 fmt.Println("Saldo Actualizado de Ana : " , WalletAna.WalletBalance)
	 fmt.Printf("\n")

	 transactionSuccefull , err := aplication.Transfer(&WalletGabriel,&WalletAna,2500)

	 if err != nil {
		fmt.Println(err)
		return
	 }

	fmt.Printf("%s : %d\n" , UserGabriel.UserName , WalletGabriel.WalletBalance)
	fmt.Printf("%s : %d\n" , UserAna.UserName , WalletAna.WalletBalance)

	fmt.Println("Transaction Succefull")
	fmt.Println(transactionSuccefull)


}