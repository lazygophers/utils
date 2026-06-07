//go:build country_all || country_americas || country_caribbean || country_ht || currency_all || currency_htg

package currency

// Htg — ISO 4217 HTG.
var Htg = New("HTG", "G", 332).
	WithDecimals(2).
	WithBanknotes(10, 25, 50, 100, 250, 500, 1000).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 5)
