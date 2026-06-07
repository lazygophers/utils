//go:build country_all || country_asia || country_ph || country_south_eastern_asia || currency_all || currency_php

package currency

// Php — ISO 4217 PHP.
var Php = New("PHP", "₱", 608).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(1, 5, 10, 25)
