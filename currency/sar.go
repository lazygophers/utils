//go:build country_all || country_asia || country_sa || country_western_asia || currency_all || currency_sar

package currency

// SAR — ISO 4217 SAR.
var SAR = New("SAR", "ر.س", 682).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 50, 100, 200, 500).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1, 2)
