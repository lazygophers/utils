//go:build country_africa || country_all || country_na || country_southern_africa || currency_all || currency_nad

package currency

// NAD — ISO 4217 NAD.
var NAD = New("NAD", "$", 516).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.5, 1, 5, 10)
