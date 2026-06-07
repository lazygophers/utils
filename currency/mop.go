//go:build country_all || country_asia || country_eastern_asia || country_mo || currency_all || currency_mop

package currency

// Mop — ISO 4217 MOP.
var Mop = New("MOP", "MOP$", 446).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 500, 1000).
	WithCoins(0.1, 0.2, 0.5, 1, 2, 5, 10)
