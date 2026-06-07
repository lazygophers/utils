//go:build country_all || country_melanesia || country_oceania || country_sb || currency_all || currency_sbd

package currency

// Sbd — ISO 4217 SBD.
var Sbd = New("SBD", "$", 90).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2)
