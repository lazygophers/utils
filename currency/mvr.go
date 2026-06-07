//go:build country_all || country_asia || country_mv || country_southern_asia || currency_all || currency_mvr

package currency

// Mvr — ISO 4217 MVR.
var Mvr = New("MVR", "Rf", 462).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 500, 1000).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1, 2)
