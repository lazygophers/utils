//go:build country_all || country_americas || country_caribbean || country_cu || currency_all || currency_cup

package currency

// CUP — ISO 4217 CUP.
var CUP = New("CUP", "$", 192).
	WithDecimals(2).
	WithBanknotes(1, 3, 5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(1, 5, 10, 20, 40, 50, 100)
