//go:build country_all || country_eastern_europe || country_europe || country_pl || currency_all || currency_pln

package currency

// Pln — ISO 4217 PLN.
var Pln = New("PLN", "zł", 985).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 500).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2, 5)
