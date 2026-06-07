//go:build country_africa || country_all || country_eastern_africa || country_rw || currency_all || currency_rwf

package currency

// RWF — ISO 4217 RWF.
var RWF = New("RWF", "R₣", 646).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000).
	WithCoins(1, 2, 5, 10, 20, 50, 100)
