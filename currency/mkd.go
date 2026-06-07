//go:build country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd

package currency

// Mkd — ISO 4217 MKD.
var Mkd = New("MKD", "ден", 807).
	WithDecimals(2).
	WithBanknotes(10, 50, 100, 200, 500, 1000, 2000, 5000).
	WithCoins(1, 2, 5, 10, 50)
