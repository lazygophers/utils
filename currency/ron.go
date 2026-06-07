//go:build country_all || country_eastern_europe || country_europe || country_ro || currency_all || currency_ron

package currency

// RON — ISO 4217 RON.
var RON = New("RON", "lei", 946).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 50, 100, 200, 500).
	WithCoins(0.01, 0.05, 0.1, 0.5)
