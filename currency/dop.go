//go:build country_all || country_americas || country_caribbean || country_do || currency_all || currency_dop

package currency

// Dop — ISO 4217 DOP.
var Dop = New("DOP", "RD$", 214).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000, 2000).
	WithCoins(1, 5, 10, 25)
