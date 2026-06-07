//go:build country_africa || country_all || country_eastern_africa || country_et || currency_all || currency_etb

package currency

// Etb — ISO 4217 ETB.
var Etb = New("ETB", "Br", 230).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 50, 100, 200).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
