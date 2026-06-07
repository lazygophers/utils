//go:build country_af || country_all || country_asia || country_southern_asia || currency_afn || currency_all

package currency

// Afn — ISO 4217 AFN.
var Afn = New("AFN", "؋", 971).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100, 500, 1000).
	WithCoins(1, 2, 5)
