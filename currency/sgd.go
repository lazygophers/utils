package currency

// Sgd — ISO 4217 SGD.
var Sgd = New("SGD", "S$", 702).
	WithDecimals(2).
	WithBanknotes(2, 5, 10, 50, 100, 1000).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1)
