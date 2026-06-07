//go:build country_all || country_americas || country_bs || country_caribbean || currency_all || currency_bsd

package currency

// Bsd — ISO 4217 BSD.
var Bsd = New("BSD", "$", 44).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.15, 0.25, 0.5, 1, 2)
