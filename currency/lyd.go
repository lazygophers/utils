//go:build country_africa || country_all || country_ly || country_northern_africa || currency_all || currency_lyd

package currency

// Lyd — ISO 4217 LYD.
var Lyd = New("LYD", "ل.د", 434).
	WithDecimals(3).
	WithBanknotes(1, 5, 10, 20, 50).
	WithCoins(0.05, 0.1, 0.25, 0.5)
