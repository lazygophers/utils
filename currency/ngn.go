//go:build country_africa || country_all || country_ng || country_western_africa || currency_all || currency_ngn

package currency

// Ngn — ISO 4217 NGN.
var Ngn = New("NGN", "₦", 566).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100, 200, 500, 1000).
	WithCoins(0.5, 1, 2)
