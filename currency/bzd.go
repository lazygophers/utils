//go:build country_all || country_americas || country_bz || country_central_america || currency_all || currency_bzd

package currency

// Bzd — ISO 4217 BZD.
var Bzd = New("BZD", "BZ$", 84).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
