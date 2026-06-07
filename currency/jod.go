//go:build country_all || country_asia || country_jo || country_western_asia || currency_all || currency_jod

package currency

// JOD — ISO 4217 JOD.
var JOD = New("JOD", "JD", 400).
	WithDecimals(3).
	WithBanknotes(1, 5, 10, 20, 50).
	WithCoins(0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1)
