//go:build country_africa || country_all || country_eastern_africa || country_ss || currency_all || currency_ssp

package currency

// Ssp — ISO 4217 SSP.
var Ssp = New("SSP", "£", 728).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 25, 50, 100)
