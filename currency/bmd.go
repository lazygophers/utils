//go:build country_all || country_americas || country_bm || country_northern_america || currency_all || currency_bmd

package currency

// BMD — ISO 4217 BMD.
var BMD = New("BMD", "$", 60).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 1)
