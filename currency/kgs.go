//go:build country_all || country_asia || country_central_asia || country_kg || currency_all || currency_kgs

package currency

// Kgs — ISO 4217 KGS.
var Kgs = New("KGS", "сом", 417).
	WithDecimals(2).
	WithBanknotes(20, 50, 100, 200, 500, 1000, 5000).
	WithCoins(1, 3, 5, 10)
