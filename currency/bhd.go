//go:build country_all || country_asia || country_bh || country_western_asia || currency_all || currency_bhd

package currency

// BHD — ISO 4217 BHD.
var BHD = New("BHD", ".د.ب", 48).
	WithDecimals(3).
	WithBanknotes(0.5, 1, 5, 10, 20).
	WithCoins(0.005, 0.01, 0.025, 0.05, 0.1)
