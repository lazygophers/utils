//go:build country_all || country_americas || country_central_america || country_gt || currency_all || currency_gtq

package currency

// Gtq — ISO 4217 GTQ.
var Gtq = New("GTQ", "Q", 320).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100, 200).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
