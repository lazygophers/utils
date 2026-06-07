//go:build country_all || country_asia || country_my || country_south_eastern_asia || currency_all || currency_myr

package currency

// MYR — ISO 4217 MYR.
var MYR = New("MYR", "RM", 458).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5)
