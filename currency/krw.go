package currency

// Krw — ISO 4217 KRW.
var Krw = New("KRW", "₩", 410).
	WithDecimals(0).
	WithBanknotes(1000, 5000, 10000, 50000).
	WithCoins(10, 50, 100, 500)
