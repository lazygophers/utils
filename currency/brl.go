//go:build country_all || country_americas || country_br || country_south_america || currency_all || currency_brl

package currency

// Brl — ISO 4217 BRL.
var Brl = New("BRL", "R$", 986).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.25, 0.5, 1)
