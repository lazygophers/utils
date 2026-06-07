//go:build country_all || country_asia || country_western_asia || country_ye || currency_all || currency_yer

package currency

// Yer — ISO 4217 YER.
var Yer = New("YER", "﷼", 886).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 250, 500, 1000).
	WithCoins(1, 5, 10, 20)
