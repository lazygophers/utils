//go:build country_africa || country_all || country_eastern_africa || country_mw || currency_all || currency_mwk

package currency

// MWK — ISO 4217 MWK.
var MWK = New("MWK", "MK", 454).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000, 2000).
	WithCoins(1, 5, 10)
