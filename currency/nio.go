//go:build country_all || country_americas || country_central_america || country_ni || currency_all || currency_nio

package currency

// Nio — ISO 4217 NIO.
var Nio = New("NIO", "C$", 558).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 500, 1000).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1, 5, 10)
