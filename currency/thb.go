//go:build country_all || country_asia || country_south_eastern_asia || country_th || currency_all || currency_thb

package currency

// THB — ISO 4217 THB.
var THB = New("THB", "฿", 764).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 500, 1000).
	WithCoins(0.25, 0.5, 1, 2, 5, 10)
