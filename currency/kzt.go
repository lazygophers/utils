//go:build country_all || country_asia || country_central_asia || country_kz || currency_all || currency_kzt

package currency

// Kzt — ISO 4217 KZT.
var Kzt = New("KZT", "₸", 398).
	WithDecimals(2).
	WithBanknotes(200, 500, 1000, 2000, 5000, 10000, 20000).
	WithCoins(1, 2, 5, 10, 20, 50, 100)
