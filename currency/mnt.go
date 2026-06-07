//go:build country_all || country_asia || country_eastern_asia || country_mn || currency_all || currency_mnt

package currency

// Mnt — ISO 4217 MNT.
var Mnt = New("MNT", "₮", 496).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 500, 1000, 5000, 10000, 20000)
