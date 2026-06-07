//go:build country_all || country_by || country_eastern_europe || country_europe || currency_all || currency_byn

package currency

// BYN — ISO 4217 BYN.
var BYN = New("BYN", "Br", 933).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200, 500).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
