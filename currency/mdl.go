//go:build country_all || country_eastern_europe || country_europe || country_md || currency_all || currency_mdl

package currency

// MDL — ISO 4217 MDL.
var MDL = New("MDL", "L", 498).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(1, 2, 5, 10, 25, 50)
