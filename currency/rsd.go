//go:build country_all || country_europe || country_rs || country_southern_europe || currency_all || currency_rsd

package currency

// RSD — ISO 4217 RSD.
var RSD = New("RSD", "din", 941).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 500, 1000, 2000, 5000).
	WithCoins(1, 2, 5, 10, 20)
