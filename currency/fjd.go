//go:build country_all || country_fj || country_melanesia || country_oceania || currency_all || currency_fjd

package currency

// FJD — ISO 4217 FJD.
var FJD = New("FJD", "$", 242).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2)
