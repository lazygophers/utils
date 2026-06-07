//go:build country_all || country_americas || country_pe || country_south_america || currency_all || currency_pen

package currency

// Pen — ISO 4217 PEN.
var Pen = New("PEN", "S/.", 604).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.1, 0.2, 0.5, 1, 2, 5)
