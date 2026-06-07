//go:build country_africa || country_all || country_bi || country_eastern_africa || currency_all || currency_bif

package currency

// Bif — ISO 4217 BIF.
var Bif = New("BIF", "FBu", 108).
	WithDecimals(0).
	WithBanknotes(10, 20, 50, 100, 500, 1000, 2000, 5000, 10000).
	WithCoins(1, 5, 10, 50)
