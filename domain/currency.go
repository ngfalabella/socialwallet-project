package domain

type Currency string

const (
	CurrencyARS Currency = "ARS"
	CurrencyUSD Currency = "USD"
	CurrencyBTC Currency = "BTC"
)

func (currency Currency) IsValidCurrency() bool {
	switch currency {
	case CurrencyARS,CurrencyUSD,CurrencyBTC :
		return true
	default :
		return false	
	}
}