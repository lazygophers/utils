//go:build country_all || country_antarctic || country_au || country_australia_and_new_zealand || country_cc || country_cx || country_hm || country_ki || country_micronesia || country_nf || country_nr || country_oceania || country_polynesia || country_tv || currency_all || currency_aud

package currency

// Aud — ISO 4217 AUD.
var Aud = New("AUD", "A$", 36).
	WithDecimals(2).
	WithBanknotes(5, 10, 20, 50, 100).
	WithCoins(0.05, 0.1, 0.2, 0.5, 1, 2).
	WithReserve()
