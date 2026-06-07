//go:build country_all || country_asia || country_mm || country_south_eastern_asia || currency_all || currency_mmk

package currency

// Mmk — ISO 4217 MMK.
var Mmk = New("MMK", "K", 104).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000, 5000, 10000)
