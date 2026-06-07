//go:build country_africa || country_all || country_cd || country_middle_africa || currency_all || currency_cdf

package currency

// CDF — ISO 4217 CDF.
var CDF = New("CDF", "FC", 976).
	WithDecimals(2).
	WithBanknotes(50, 100, 200, 500, 1000, 5000, 10000, 20000)
