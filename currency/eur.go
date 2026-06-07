package currency

// Eur — ISO 4217 EUR.
var Eur = New("EUR", "€", 978).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200, 500).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2).
	WithReserve()
