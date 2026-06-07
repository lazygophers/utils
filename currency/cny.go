package currency

// CNY — ISO 4217 CNY.
var CNY = New("CNY", "¥", 156).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(0.1, 0.5, 1).
	WithReserve()
