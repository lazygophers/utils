//go:build country_africa || country_all || country_eastern_africa || country_mu || currency_all || currency_mur

package currency

// Mur — ISO 4217 MUR.
var Mur = New("MUR", "₨", 480).
	WithDecimals(2).
	WithBanknotes(25, 50, 100, 200, 500, 1000, 2000).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 5, 10, 20)
