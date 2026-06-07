//go:build country_africa || country_all || country_eastern_africa || country_mg || currency_all || currency_mga

package currency

// MGA — ISO 4217 MGA.
var MGA = New("MGA", "Ar", 969).
	WithDecimals(0).
	WithBanknotes(100, 200, 500, 1000, 2000, 5000, 10000, 20000).
	WithCoins(1, 2, 5, 10, 20, 50)
