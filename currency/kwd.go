//go:build country_all || country_asia || country_kw || country_western_asia || currency_all || currency_kwd

package currency

// Kwd — ISO 4217 KWD.
var Kwd = New("KWD", "KD", 414).
	WithDecimals(3).
	WithBanknotes(0.25, 0.5, 1, 5, 10, 20).
	WithCoins(0.005, 0.01, 0.02, 0.05, 0.1)
