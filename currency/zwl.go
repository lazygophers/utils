//go:build country_africa || country_all || country_eastern_africa || country_zw || currency_all || currency_zwl

package currency

// Zwl — ISO 4217 ZWL.
var Zwl = New("ZWL", "$", 932).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100, 500).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
