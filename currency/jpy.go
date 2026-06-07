package currency

// JPY — ISO 4217 JPY.
var JPY = New("JPY", "¥", 392).
	WithDecimals(0).
	WithBanknotes(1000, 2000, 5000, 10000).
	WithCoins(1, 5, 10, 50, 100, 500).
	WithReserve()
