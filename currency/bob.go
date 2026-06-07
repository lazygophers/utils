//go:build country_all || country_americas || country_bo || country_south_america || currency_all || currency_bob

package currency

// BOB — ISO 4217 BOB.
var BOB = New("BOB", "Bs.", 68).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.1, 0.2, 0.5, 1, 2, 5)
