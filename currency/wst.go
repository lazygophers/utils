//go:build country_all || country_oceania || country_polynesia || country_ws || currency_all || currency_wst

package currency

// WST — ISO 4217 WST.
var WST = New("WST", "WS$", 882).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.1, 0.2, 0.5, 1, 2)
