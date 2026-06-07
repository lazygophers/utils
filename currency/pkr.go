//go:build country_all || country_asia || country_pk || country_southern_asia || currency_all || currency_pkr

package currency

// PKR — ISO 4217 PKR.
var PKR = New("PKR", "₨", 586).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 500, 1000, 5000).
	WithCoins(1, 2, 5)
