package domain

type Currency string

const (
	CurrencyARS Currency = "ARS"
	CurrencyUSD Currency = "USD"
	CurrencyBTC Currency = "BTC"
)

func ( currency Currency) IsValid() bool {
	switch currency {
	case CurrencyARS , CurrencyUSD , CurrencyBTC :
		return true
	default : 
		return false	
	}	
}