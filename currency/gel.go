//go:build country_all || country_asia || country_ge || country_western_asia || currency_all || currency_gel

package currency

// GEL — ISO 4217 GEL.
var GEL = New("GEL", "₾", 981).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
