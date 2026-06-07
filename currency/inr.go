package currency

// Inr — ISO 4217 INR.
var Inr = New("INR", "₹", 356).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200, 500, 2000).
	WithCoins(1, 2, 5, 10, 20)
