//go:build country_all || country_asia || country_central_asia || country_tj || currency_all || currency_tjs

package currency

// TJS — ISO 4217 TJS.
var TJS = New("TJS", "ЅМ", 972).
	WithDecimals(2).
	WithBanknotes(1, 3, 5, 10, 20, 50, 100, 200, 500).
	WithCoins(0.05, 0.1, 0.2, 0.25, 0.5, 1, 2, 3, 5)
