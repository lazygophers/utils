//go:build country_africa || country_all || country_mr || country_western_africa || currency_all || currency_mru

package currency

// Mru — ISO 4217 MRU.
var Mru = New("MRU", "UM", 929).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(0.2, 0.5, 1, 5, 10, 20, 50)
