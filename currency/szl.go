//go:build country_africa || country_all || country_southern_africa || country_sz || currency_all || currency_szl

package currency

// Szl — ISO 4217 SZL.
var Szl = New("SZL", "L", 748).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5)
