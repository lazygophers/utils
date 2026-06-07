//go:build country_all || country_oceania || country_polynesia || country_to || currency_all || currency_top

package currency

// TOP — ISO 4217 TOP.
var TOP = New("TOP", "T$", 776).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2)
