//go:build country_all || country_bg || country_eastern_europe || country_europe || currency_all || currency_bgn

package currency

// Bgn — ISO 4217 BGN.
var Bgn = New("BGN", "лв", 975).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
