//go:build country_africa || country_all || country_eg || country_northern_africa || currency_all || currency_egp

package currency

// EGP — ISO 4217 EGP.
var EGP = New("EGP", "£", 818).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200).
	WithCoins(0.25, 0.5, 1)
