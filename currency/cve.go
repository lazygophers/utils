//go:build country_africa || country_all || country_cv || country_western_africa || currency_all || currency_cve

package currency

// CVE — ISO 4217 CVE.
var CVE = New("CVE", "$", 132).
	WithDecimals(2).
	WithBanknotes(200, 500, 1000, 2000, 5000).
	WithCoins(1, 5, 10, 20, 50, 100, 200)
