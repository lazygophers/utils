//go:build country_all || country_asia || country_il || country_ps || country_western_asia || currency_all || currency_ils

package currency

// Ils — ISO 4217 ILS.
var Ils = New("ILS", "₪", 376).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200).
	WithCoins(0.1, 0.5, 1, 2, 5, 10)
