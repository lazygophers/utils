//go:build country_africa || country_all || country_sh || country_western_africa || currency_all || currency_shp

package currency

// Shp — ISO 4217 SHP.
var Shp = New("SHP", "£", 654).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
