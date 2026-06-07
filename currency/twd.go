package currency

// TWD — ISO 4217 TWD.
var TWD = New("TWD", "NT$", 901).
	WithDecimals(2).
	WithBanknotes(100, 200, 500, 1000, 2000).
	WithCoins(1, 5, 10, 20, 50)
