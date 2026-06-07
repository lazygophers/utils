//go:build country_africa || country_all || country_sl || country_western_africa || currency_all || currency_sll

package currency

// SLL — ISO 4217 SLL.
var SLL = New("SLL", "Le", 694).
	WithDecimals(2).
	WithBanknotes(1000, 2000, 5000, 10000).
	WithCoins(10, 50, 100, 500)
