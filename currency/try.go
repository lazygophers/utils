//go:build country_all || country_asia || country_tr || country_western_asia || currency_all || currency_try

package currency

// Try — ISO 4217 TRY.
var Try = New("TRY", "₺", 949).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200).
	WithCoins(0.01, 0.05, 0.1, 0.25, 0.5, 1)
