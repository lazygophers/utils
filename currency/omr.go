//go:build country_all || country_asia || country_om || country_western_asia || currency_all || currency_omr

package currency

// Omr — ISO 4217 OMR.
var Omr = New("OMR", "ر.ع.", 512).
	WithDecimals(3).
	WithBanknotes(0.1, 0.2, 0.5, 1, 5, 10, 20, 50).
	WithCoins(0.005, 0.01, 0.025, 0.05, 0.1)
