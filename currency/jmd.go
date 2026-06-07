//go:build country_all || country_americas || country_caribbean || country_jm || currency_all || currency_jmd

package currency

// JMD — ISO 4217 JMD.
var JMD = New("JMD", "J$", 388).
	WithDecimals(2).
	WithBanknotes(50, 100, 500, 1000, 5000).
	WithCoins(1, 5, 10, 20)
