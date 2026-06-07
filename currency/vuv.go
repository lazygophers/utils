//go:build country_all || country_melanesia || country_oceania || country_vu || currency_all || currency_vuv

package currency

// VUV — ISO 4217 VUV.
var VUV = New("VUV", "VT", 548).
	WithDecimals(0).
	WithBanknotes(200, 500, 1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 20, 50, 100)
