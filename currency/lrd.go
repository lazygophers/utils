//go:build country_africa || country_all || country_lr || country_western_africa || currency_all || currency_lrd

package currency

// LRD — ISO 4217 LRD.
var LRD = New("LRD", "$", 430).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 500, 1000).
	WithCoins(5, 10, 25, 50)
