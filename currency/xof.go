//go:build country_africa || country_all || country_bf || country_bj || country_ci || country_gw || country_ml || country_ne || country_sn || country_tg || country_western_africa || currency_all || currency_xof

package currency

// Xof — ISO 4217 XOF.
var Xof = New("XOF", "CFA", 952).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 25, 50, 100, 200, 250, 500)
