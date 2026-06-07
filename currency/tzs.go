//go:build country_africa || country_all || country_eastern_africa || country_tz || currency_all || currency_tzs

package currency

// Tzs — ISO 4217 TZS.
var Tzs = New("TZS", "TSh", 834).
	WithDecimals(2).
	WithBanknotes(500, 1000, 2000, 5000, 10000).
	WithCoins(50, 100, 200, 500)
