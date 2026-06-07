//go:build country_africa || country_all || country_northern_africa || country_sd || currency_all || currency_sdg

package currency

// Sdg — ISO 4217 SDG.
var Sdg = New("SDG", "ج.س.", 938).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100, 200, 500).
	WithCoins(0.1, 0.25, 0.5, 1)
