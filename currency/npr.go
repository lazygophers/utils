//go:build country_all || country_asia || country_np || country_southern_asia || currency_all || currency_npr

package currency

// NPR — ISO 4217 NPR.
var NPR = New("NPR", "रू", 524).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 25, 50, 100, 500, 1000).
	WithCoins(1, 2)
