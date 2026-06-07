//go:build country_all || country_americas || country_south_america || country_sr || currency_all || currency_srd

package currency

// Srd — ISO 4217 SRD.
var Srd = New("SRD", "$", 968).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 1, 5, 10, 25, 50, 100, 250)
