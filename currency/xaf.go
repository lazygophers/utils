//go:build country_africa || country_all || country_cf || country_cg || country_cm || country_ga || country_gq || country_middle_africa || country_td || currency_all || currency_xaf

package currency

// Xaf — ISO 4217 XAF.
var Xaf = New("XAF", "FCFA", 950).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 25, 50, 100, 500)
