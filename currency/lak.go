//go:build country_all || country_asia || country_la || country_south_eastern_asia || currency_all || currency_lak

package currency

// LAK — ISO 4217 LAK.
var LAK = New("LAK", "₭", 418).
	WithDecimals(2).
	WithBanknotes(500, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
