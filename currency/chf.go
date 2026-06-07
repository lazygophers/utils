//go:build country_all || country_ch || country_europe || country_li || country_western_europe || currency_all || currency_chf

package currency

// CHF — ISO 4217 CHF.
var CHF = New("CHF", "Fr", 756).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 1000).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5).
	WithReserve()
