//go:build country_all || country_asia || country_az || country_western_asia || currency_all || currency_azn

package currency

// Azn — ISO 4217 AZN.
var Azn = New("AZN", "₼", 944).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100, 200).
	WithCoins(0.01, 0.03, 0.05, 0.1, 0.2, 0.5)
