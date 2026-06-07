//go:build country_all || country_americas || country_central_america || country_pa || currency_all || currency_pab

package currency

// Pab — ISO 4217 PAB.
var Pab = New("PAB", "B/.", 590).
	WithDecimals(2).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
