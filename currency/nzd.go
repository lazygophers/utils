//go:build country_all || country_australia_and_new_zealand || country_ck || country_nu || country_nz || country_oceania || country_pn || country_polynesia || country_tk || currency_all || currency_nzd

package currency

// Nzd — ISO 4217 NZD.
var Nzd = New("NZD", "$", 554).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.1, 0.2, 0.5, 1, 2)
