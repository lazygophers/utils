//go:build country_all || country_americas || country_co || country_south_america || currency_all || currency_cop

package currency

// COP — ISO 4217 COP.
var COP = New("COP", "$", 170).
	WithDecimals(2).
	WithBanknotes(2000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(50, 100, 200, 500, 1000)
