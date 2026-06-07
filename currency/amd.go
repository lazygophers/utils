//go:build country_all || country_am || country_asia || country_western_asia || currency_all || currency_amd

package currency

// Amd — ISO 4217 AMD.
var Amd = New("AMD", "֏", 51).
	WithDecimals(2).
	WithBanknotes(500, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(10, 20, 50, 100, 200, 500)
