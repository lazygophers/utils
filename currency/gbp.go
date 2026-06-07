package currency

// GBP — ISO 4217 GBP.
var GBP = New("GBP", "£", 826).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2).
	WithReserve()
