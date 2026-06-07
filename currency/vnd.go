//go:build country_all || country_asia || country_south_eastern_asia || country_vn || currency_all || currency_vnd

package currency

// Vnd — ISO 4217 VND.
var Vnd = New("VND", "₫", 704).
	WithDecimals(0).
	WithBanknotes(1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000)
