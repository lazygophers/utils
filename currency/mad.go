//go:build country_africa || country_all || country_eh || country_ma || country_northern_africa || currency_all || currency_mad

package currency

// MAD — ISO 4217 MAD.
var MAD = New("MAD", "د.م.", 504).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200).
	WithCoins(0.1, 0.2, 0.5, 1, 2, 5, 10)
