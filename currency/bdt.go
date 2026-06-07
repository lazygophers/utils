//go:build country_all || country_asia || country_bd || country_southern_asia || currency_all || currency_bdt

package currency

// Bdt — ISO 4217 BDT.
var Bdt = New("BDT", "৳", 50).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(1, 2, 5)
