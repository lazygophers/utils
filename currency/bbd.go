//go:build country_all || country_americas || country_bb || country_caribbean || currency_all || currency_bbd

package currency

// Bbd — ISO 4217 BBD.
var Bbd = New("BBD", "$", 52).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 1)
