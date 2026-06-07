//go:build country_africa || country_all || country_eastern_africa || country_zm || currency_all || currency_zmw

package currency

// ZMW — ISO 4217 ZMW.
var ZMW = New("ZMW", "ZK", 967).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.5, 1)
