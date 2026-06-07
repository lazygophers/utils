//go:build country_africa || country_all || country_eastern_africa || country_so || currency_all || currency_sos

package currency

// Sos — ISO 4217 SOS.
var Sos = New("SOS", "S", 706).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000).
	WithCoins(0.05, 0.1, 0.5, 1)
