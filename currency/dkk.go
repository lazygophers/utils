//go:build country_all || country_americas || country_dk || country_europe || country_fo || country_gl || country_northern_america || country_northern_europe || currency_all || currency_dkk

package currency

// DKK — ISO 4217 DKK.
var DKK = New("DKK", "kr", 208).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000).
	WithCoins(0.5, 1, 2, 5, 10, 20)
