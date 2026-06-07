//go:build country_all || country_europe || country_gi || country_southern_europe || currency_all || currency_gip

package currency

// GIP — ISO 4217 GIP.
var GIP = New("GIP", "£", 292).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1, 2)
