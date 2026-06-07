//go:build country_all || country_asia || country_kh || country_south_eastern_asia || currency_all || currency_khr

package currency

// Khr — ISO 4217 KHR.
var Khr = New("KHR", "៛", 116).
	WithDecimals(2).
	WithBanknotes(100, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
