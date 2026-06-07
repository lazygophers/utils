//go:build country_all || country_asia || country_eastern_asia || country_kp || currency_all || currency_kpw

package currency

// Kpw — ISO 4217 KPW.
var Kpw = New("KPW", "₩", 408).
	WithDecimals(2).
	WithBanknotes(5, 10, 50, 100, 200, 500, 1000, 2000, 5000).
	WithCoins(1, 5, 10, 50)
