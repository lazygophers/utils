//go:build country_all || country_asia || country_central_asia || country_tm || currency_all || currency_tmt

package currency

// Tmt — ISO 4217 TMT.
var Tmt = New("TMT", "T", 934).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
