//go:build country_all || country_europe || country_is || country_northern_europe || currency_all || currency_isk

package currency

// Isk — ISO 4217 ISK.
var Isk = New("ISK", "kr", 352).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000, 10000).
	WithCoins(1, 5, 10, 50, 100)
