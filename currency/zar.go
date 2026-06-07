//go:build country_africa || country_all || country_southern_africa || country_za || currency_all || currency_zar

package currency

// Zar — ISO 4217 ZAR.
var Zar = New("ZAR", "R", 710).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5)
