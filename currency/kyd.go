//go:build country_all || country_americas || country_caribbean || country_ky || currency_all || currency_kyd

package currency

// KYD — ISO 4217 KYD.
var KYD = New("KYD", "$", 136).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 25, 50, 100).
	WithCoins(0.01, 0.05, 0.1, 0.25)
