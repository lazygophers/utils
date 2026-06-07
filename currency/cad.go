//go:build country_all || country_americas || country_ca || country_northern_america || currency_all || currency_cad

package currency

// CAD — ISO 4217 CAD.
var CAD = New("CAD", "C$", 124).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.25, 1, 2).
	WithReserve()
