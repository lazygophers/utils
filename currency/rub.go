package currency

// RUB — ISO 4217 RUB.
var RUB = New("RUB", "₽", 643).
	WithDecimals(2).
	WithBanknotes(5, 10, 50, 100, 200, 500, 1000, 2000, 5000).
	WithCoins(0.01, 0.05, 0.1, 0.5, 1, 2, 5, 10)
