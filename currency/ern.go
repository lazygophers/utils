//go:build country_africa || country_all || country_eastern_africa || country_er || currency_all || currency_ern

package currency

// Ern — ISO 4217 ERN.
var Ern = New("ERN", "Nfk", 232).
	WithDecimals(2).
	WithBanknotes(1, 5, 10, 20, 50, 100).
	WithCoins(1, 5, 10, 25, 50, 100)
