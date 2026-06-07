//go:build country_all || country_asia || country_qa || country_western_asia || currency_all || currency_qar

package currency

// QAR — ISO 4217 QAR.
var QAR = New("QAR", "ر.ق", 634).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 50, 100, 500).
	WithCoins(0.25, 0.5, 1)
