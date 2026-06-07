//go:build country_all || country_asia || country_sy || country_western_asia || currency_all || currency_syp

package currency

// Syp — ISO 4217 SYP.
var Syp = New("SYP", "£S", 760).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000, 2000, 5000).
	WithCoins(1, 2, 5, 10, 25)
