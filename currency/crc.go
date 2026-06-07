//go:build country_all || country_americas || country_central_america || country_cr || currency_all || currency_crc

package currency

// CRC — ISO 4217 CRC.
var CRC = New("CRC", "₡", 188).
	WithDecimals(2).
	WithBanknotes(1000, 2000, 5000, 10000, 20000, 50000).
	WithCoins(5, 10, 25, 50, 100, 500)
