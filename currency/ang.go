//go:build country_all || country_americas || country_caribbean || country_cw || country_sx || currency_all || currency_ang

package currency

// ANG — ISO 4217 ANG.
var ANG = New("ANG", "ƒ", 532).
	WithDecimals(2).
	WithBanknotes(10, 25, 50, 100, 250).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5)
