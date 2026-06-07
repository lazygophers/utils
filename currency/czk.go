//go:build country_all || country_cz || country_eastern_europe || country_europe || currency_all || currency_czk

package currency

// Czk — ISO 4217 CZK.
var Czk = New("CZK", "Kč", 203).
	WithDecimals(2).
	WithBanknotes(100, 200, 500, 1000, 2000, 5000).
	WithCoins(1, 2, 5, 10, 20, 50)
