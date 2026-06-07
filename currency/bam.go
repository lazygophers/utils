//go:build country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam

package currency

// BAM — ISO 4217 BAM.
var BAM = New("BAM", "KM", 977).
	WithDecimals(2).
	WithBanknotes(10, 20, 50, 100, 200).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2, 5)
