//go:build country_all || country_asia || country_bt || country_southern_asia || currency_all || currency_btn

package currency

// Btn — ISO 4217 BTN.
var Btn = New("BTN", "Nu.", 64).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100, 500, 1000).
	WithCoins(0.2, 0.25, 0.5, 1)
