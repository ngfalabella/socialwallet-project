package domain

import "testing"

func TestValidOperation(t *testing.T) {
	newWalletCreated , err := NewWallet(1 , CurrencyARS)

	if err != nil {
		t.Fatalf("Ocurrio un error al crear wallet : %v" , 
		err ,
		)
	}

	err = newWalletCreated.Deposit(1000)

	if err != nil {
		t.Fatalf("Ocurrio un error al Depositar : %v" , err)
	}

	if newWalletCreated.WalletBalance != 1000 {
		t.Fatalf("Se esperaba saldo de $500 y se obtuvo $%d" , newWalletCreated.WalletBalance)
	}

	err = newWalletCreated.Withdraw(300)

	if err != nil {
		t.Fatalf("Ocurrio un error al hacer la extraccion : %v" , err)
	}

	if newWalletCreated.WalletBalance != 700 {
		t.Fatalf("Se esperaba un saldo de 700 y se obtuvo :%v" , newWalletCreated.WalletBalance)
	}

	err = newWalletCreated.Deposit(200)

	if err != nil {
		t.Fatalf("Ocurrio un fallo al hacer el segundo deposito : %v" , err)
	}

	if newWalletCreated.WalletBalance != 900 {
		t.Fatalf("Se esperaba un saldo de $900 y se obtuvo $%d" , newWalletCreated.WalletBalance)
	}
}

func TestExtractAllMount( t *testing.T) {
	newWalletCreated , err := NewWallet(2 , CurrencyARS)

	if err != nil {
		t.Fatalf("Ocurrio un error al crear wallet : %v" , err)
	}

	err = newWalletCreated.Deposit(700)

	if err != nil {
		t.Fatalf("Ocurrio un error al hacer deposito : %v" , err)
	}

	if newWalletCreated.WalletBalance != 700 {
		t.Fatalf("Se esperaba tener 700 y se obtuvo %d" , newWalletCreated.WalletBalance)
	}

	err = newWalletCreated.Withdraw(700)

	if err != nil {
		t.Fatalf("Ocurrio un error al hacer la extraccion %v" , err)
	}

	if newWalletCreated.WalletBalance != 0 {
		t.Fatalf("El cliente no deberia tener saldo y termino teniendo : %d" , newWalletCreated.WalletBalance)
	}
}

func TestSaldoInsuficiente ( t *testing.T) {
	newWalletCreated , err := NewWallet(3 , CurrencyARS)

	if err != nil {
		t.Fatalf("Error al crear wallet : %v" , err)
	}

	err = newWalletCreated.Deposit(400)

	if err != nil {
		t.Fatalf("Ocurrio un error al depositar : %v" ,err)
	}

	if newWalletCreated.WalletBalance != 400 {
		t.Fatalf("Se esperaba tener $400 y hay  %d " , newWalletCreated.WalletBalance)
	}

	err =newWalletCreated.Withdraw(500)

	e

}