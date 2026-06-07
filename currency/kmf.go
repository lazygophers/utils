//go:build country_africa || country_all || country_eastern_africa || country_km || currency_all || currency_kmf

package currency

// Kmf — ISO 4217 KMF.
var Kmf = New("KMF", "CF", 174).
	WithDecimals(0).
	WithBanknotes(500, 1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 25, 50, 100, 250)
