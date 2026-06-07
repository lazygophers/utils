//go:build country_africa || country_all || country_dj || country_eastern_africa || currency_all || currency_djf

package currency

// DJF — ISO 4217 DJF.
var DJF = New("DJF", "Fdj", 262).
	WithDecimals(0).
	WithBanknotes(1000, 2000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 20, 50, 100, 250, 500)
