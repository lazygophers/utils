//go:build country_africa || country_all || country_eastern_africa || country_ke || currency_all || currency_kes

package currency

// Kes — ISO 4217 KES.
var Kes = New("KES", "KSh", 404).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000).
	WithCoins(1, 5, 10, 20)
