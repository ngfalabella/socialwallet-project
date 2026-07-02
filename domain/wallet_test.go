package domain 

import ( 
	"testing"
	"errors"
)

func TestDepositNegativeAmount(t *testing.T) {
	wallet ,err := NewWallet(1 , CurrencyARS)
	if err != nil {
		t.Fatalf("no se pudo crear la wallet : %v" , err)
	}

	err = wallet.Deposit(500)
	if err != nil {
		t.Fatalf("el deposito devolvio un error inesperado : %v" , err)
	}

	if wallet.WalletBalance != 500 {
		t.Fatalf("Saldo esperado : 500 , saldo obtenido :%d" ,wallet.WalletBalance)
	}

}

func TestWalletValidCurrency( t *testing.T) {
	wallet,err := NewWallet( 2 , CurrencyARS)
	 
	if err != nil {
		t.Fatalf("No se pudo crear wallet  :%v" , err)
	}

	err = wallet.Deposit(-500)

	if err == nil {
		t.Fatalf("Se esperaba un error pero se obtuvo nil")
	}

	if !errors.Is(err,InvalidMountWallet) {
		t.Fatalf("Se esperaba invaiidMounwaller, pero se obtuvo %v" ,err)
	}

	if wallet.WalletBalance != 0 {
	t.Fatalf(
		"el saldo debía permanecer en 0, pero quedó en %d",
		wallet.WalletBalance,
	)
}
}