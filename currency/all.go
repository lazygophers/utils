//go:build country_al || country_all || country_europe || country_southern_europe || currency_all

package currency

// ALL — ISO 4217 ALL.
var ALL = New("ALL", "L", 8).
	WithDecimals(2).
	WithBanknotes(200, 500, 1000, 2000, 5000, 10000).
	WithCoins(1, 5, 10, 20, 50, 100)
