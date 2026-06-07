//go:build country_all || country_eastern_europe || country_europe || country_hu || currency_all || currency_huf

package currency

// Huf — ISO 4217 HUF.
var Huf = New("HUF", "Ft", 348).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000, 10000, 20000).
	WithCoins(5, 10, 20, 50, 100, 200)
