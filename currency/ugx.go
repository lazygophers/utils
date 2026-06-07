//go:build country_africa || country_all || country_eastern_africa || country_ug || currency_all || currency_ugx

package currency

// UGX — ISO 4217 UGX.
var UGX = New("UGX", "USh", 800).
	WithDecimals(0).
	WithBanknotes(1000, 2000, 5000, 10000, 20000, 50000).
	WithCoins(50, 100, 200, 500, 1000)
