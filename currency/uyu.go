//go:build country_all || country_americas || country_south_america || country_uy || currency_all || currency_uyu

package currency

// Uyu — ISO 4217 UYU.
var Uyu = New("UYU", "$U", 858).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000, 2000).
	WithCoins(1, 2, 5, 10, 50)
