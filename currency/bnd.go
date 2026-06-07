//go:build country_all || country_asia || country_bn || country_south_eastern_asia || currency_all || currency_bnd

package currency

// Bnd — ISO 4217 BND.
var Bnd = New("BND", "$", 96).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100, 500, 1000, 10000).
	WithCoins(0.01, 0.05, 0.1, 0.2, 0.5, 1)
