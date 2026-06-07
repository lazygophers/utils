//go:build country_all || country_asia || country_lk || country_southern_asia || currency_all || currency_lkr

package currency

// LKR — ISO 4217 LKR.
var LKR = New("LKR", "Rs", 144).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 500, 1000, 5000).
	WithCoins(1, 2, 5, 10)
