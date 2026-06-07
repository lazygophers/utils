//go:build country_all || country_antarctic || country_bv || country_europe || country_no || country_northern_europe || country_sj || currency_all || currency_nok

package currency

// NOK — ISO 4217 NOK.
var NOK = New("NOK", "kr", 578).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000).
	WithCoins(1, 5, 10, 20)
