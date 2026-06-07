//go:build country_all || country_americas || country_central_america || country_hn || currency_all || currency_hnl

package currency

// HNL — ISO 4217 HNL.
var HNL = New("HNL", "L", 340).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100, 500).
	WithCoins(0.05, 0.1, 0.2, 0.5)
