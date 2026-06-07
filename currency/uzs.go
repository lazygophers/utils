//go:build country_all || country_asia || country_central_asia || country_uz || currency_all || currency_uzs

package currency

// Uzs — ISO 4217 UZS.
var Uzs = New("UZS", "сўм", 860).
	WithDecimals(2).
	WithBanknotes(1000, 5000, 10000, 50000, 100000).
	WithCoins(50, 100, 200, 500)
