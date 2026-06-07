//go:build country_africa || country_all || country_dz || country_northern_africa || currency_all || currency_dzd

package currency

// DZD — ISO 4217 DZD.
var DZD = New("DZD", "دج", 12).
	WithDecimals(2).
	WithBanknotes(200, 500, 1000, 2000).
	WithCoins(1, 2, 5, 10, 20, 50, 100, 200)
