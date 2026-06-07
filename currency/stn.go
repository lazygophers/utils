//go:build country_africa || country_all || country_middle_africa || country_st || currency_all || currency_stn

package currency

// STN — ISO 4217 STN.
var STN = New("STN", "Db", 930).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200).
	WithCoins(0.1, 0.2, 0.5, 1, 2)
