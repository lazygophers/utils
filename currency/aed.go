//go:build country_ae || country_all || country_asia || country_western_asia || currency_aed || currency_all

package currency

// Aed — ISO 4217 AED.
var Aed = New("AED", "د.إ", 784).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(0.25, 0.5, 1)
