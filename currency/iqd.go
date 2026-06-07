//go:build country_all || country_asia || country_iq || country_western_asia || currency_all || currency_iqd

package currency

// IQD — ISO 4217 IQD.
var IQD = New("IQD", "ع.د", 368).
	WithDecimals(3).
	WithBanknotes(250, 500, 1000, 5000, 10000, 25000, 50000)
