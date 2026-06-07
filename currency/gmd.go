//go:build country_africa || country_all || country_gm || country_western_africa || currency_all || currency_gmd

package currency

// Gmd — ISO 4217 GMD.
var Gmd = New("GMD", "D", 270).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200).
	WithCoins(1)
