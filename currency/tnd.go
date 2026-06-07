//go:build country_africa || country_all || country_northern_africa || country_tn || currency_all || currency_tnd

package currency

// Tnd — ISO 4217 TND.
var Tnd = New("TND", "د.ت", 788).
	WithDecimals(3).
	WithBanknotes(5, 10, 20, 50).
	WithCoins(0.005, 0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2, 5)
