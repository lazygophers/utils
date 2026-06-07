//go:build country_all || country_americas || country_cl || country_south_america || currency_all || currency_clp

package currency

// Clp — ISO 4217 CLP.
var Clp = New("CLP", "$", 152).
	WithDecimals(0).
	WithBanknotes(1000, 2000, 5000, 10000, 20000).
	WithCoins(10, 50, 100, 500)
