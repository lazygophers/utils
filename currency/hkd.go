package currency

// Hkd — ISO 4217 HKD.
var Hkd = New("HKD", "HK$", 344).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 500, 1000).
	WithCoins(0.1, 0.2, 0.5, 1, 2, 5, 10)
