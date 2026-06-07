//go:build country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn

package currency

// MZN — ISO 4217 MZN.
var MZN = New("MZN", "MT", 943).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(1, 2, 5, 10)
