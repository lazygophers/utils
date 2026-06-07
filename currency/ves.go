//go:build country_all || country_americas || country_south_america || country_ve || currency_all || currency_ves

package currency

// VES — ISO 4217 VES.
var VES = New("VES", "Bs.S", 928).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200, 500, 1000)
