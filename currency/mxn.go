//go:build country_all || country_americas || country_central_america || country_mx || currency_all || currency_mxn

package currency

// Mxn — ISO 4217 MXN.
var Mxn = New("MXN", "$", 484).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 20)
