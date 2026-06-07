//go:build country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd

package currency

// TTD — ISO 4217 TTD.
var TTD = New("TTD", "TT$", 780).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
