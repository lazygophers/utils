//go:build country_all || country_asia || country_ir || country_southern_asia || currency_all || currency_irr

package currency

// Irr — ISO 4217 IRR.
var Irr = New("IRR", "﷼", 364).
	WithDecimals(2).
	WithBanknotes(100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(50, 100, 250, 500, 1000)
