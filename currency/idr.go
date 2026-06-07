//go:build country_all || country_asia || country_id || country_south_eastern_asia || currency_all || currency_idr

package currency

// IDR — ISO 4217 IDR.
var IDR = New("IDR", "Rp", 360).
	WithDecimals(2).
	WithBanknotes(1000, 2000, 5000, 10000, 20000, 50000, 100000).
	WithCoins(50, 100, 200, 500, 1000)
