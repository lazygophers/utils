//go:build country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp

package currency

// FKP — ISO 4217 FKP.
var FKP = New("FKP", "£", 238).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
