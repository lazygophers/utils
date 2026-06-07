//go:build country_all || country_asia || country_lb || country_western_asia || currency_all || currency_lbp

package currency

// Lbp — ISO 4217 LBP.
var Lbp = New("LBP", "ل.ل", 422).
	WithDecimals(2).
	WithBanknotes(1000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(50, 100, 250, 500)
