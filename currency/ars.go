//go:build country_all || country_americas || country_ar || country_south_america || currency_all || currency_ars

package currency

// Ars — ISO 4217 ARS.
var Ars = New("ARS", "$", 32).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 500, 1000, 2000).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1, 2, 5, 10)
