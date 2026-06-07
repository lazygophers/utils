//go:build country_africa || country_all || country_bw || country_southern_africa || currency_all || currency_bwp

package currency

// BWP — ISO 4217 BWP.
var BWP = New("BWP", "P", 72).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1, 2, 5)
