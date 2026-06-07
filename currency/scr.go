//go:build country_africa || country_all || country_eastern_africa || country_sc || currency_all || currency_scr

package currency

// SCR — ISO 4217 SCR.
var SCR = New("SCR", "₨", 690).
	WithDecimals(2).
	WithBanknotes(25, 50, 100, 500).
	WithCoins(1, 5, 10, 25)
