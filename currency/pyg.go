//go:build country_all || country_americas || country_py || country_south_america || currency_all || currency_pyg

package currency

// PYG — ISO 4217 PYG.
var PYG = New("PYG", "₲", 600).
	WithDecimals(0).
	WithBanknotes(2000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(50, 100, 500, 1000)
