//go:build country_africa || country_all || country_gn || country_western_africa || currency_all || currency_gnf

package currency

// Gnf — ISO 4217 GNF.
var Gnf = New("GNF", "FG", 324).
	WithDecimals(0).
	WithBanknotes(100, 500, 1000, 2000, 5000, 10000, 20000).
	WithCoins(1, 5, 10, 25, 50)
