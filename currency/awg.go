//go:build country_all || country_americas || country_aw || country_caribbean || currency_all || currency_awg

package currency

// AWG — ISO 4217 AWG.
var AWG = New("AWG", "ƒ", 533).
	WithDecimals(2).
	WithBanknotes(10, 25, 50, 100, 500).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1, 2.5, 5)
