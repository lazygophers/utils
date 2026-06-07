//go:build country_all || country_europe || country_northern_europe || country_se || currency_all || currency_sek

package currency

// SEK — ISO 4217 SEK.
var SEK = New("SEK", "kr", 752).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(1, 2, 5, 10)
