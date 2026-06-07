//go:build country_all || country_eastern_europe || country_europe || country_ua || currency_all || currency_uah

package currency

// UAH — ISO 4217 UAH.
var UAH = New("UAH", "₴", 980).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(0.1, 0.5, 1, 2, 5, 10)
