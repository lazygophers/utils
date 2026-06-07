//go:build country_all || country_melanesia || country_oceania || country_pg || currency_all || currency_pgk

package currency

// PGK — ISO 4217 PGK.
var PGK = New("PGK", "K", 598).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2)
