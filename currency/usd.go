package currency

// USD — ISO 4217 USD.
var USD = New("USD", "$", 840).
	WithDecimals(2).
	WithBanknotes(1, 2, 5, 10, 20, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1).
	WithReserve()
