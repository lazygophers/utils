//go:build country_africa || country_all || country_ls || country_southern_africa || currency_all || currency_lsl

package currency

// LSL — ISO 4217 LSL.
var LSL = New("LSL", "L", 426).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5)
