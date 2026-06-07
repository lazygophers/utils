//go:build country_all || country_americas || country_gy || country_south_america || currency_all || currency_gyd

package currency

// GYD — ISO 4217 GYD.
var GYD = New("GYD", "$", 328).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 500, 1000, 5000).
	WithCoins(1, 5, 10)
