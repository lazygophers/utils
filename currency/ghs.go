//go:build country_africa || country_all || country_gh || country_western_africa || currency_all || currency_ghs

package currency

// GHS — ISO 4217 GHS.
var GHS = New("GHS", "₵", 936).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100, 200).
	WithCoins(0.01, 0.05, 0.1, 0.2, 0.5, 1, 2)
