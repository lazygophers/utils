//go:build country_africa || country_all || country_ao || country_middle_africa || currency_all || currency_aoa

package currency

// AOA — ISO 4217 AOA.
var AOA = New("AOA", "Kz", 973).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 20, 50, 100)
