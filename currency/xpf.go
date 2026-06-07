//go:build country_all || country_melanesia || country_nc || country_oceania || country_pf || country_polynesia || country_wf || currency_all || currency_xpf

package currency

// Xpf — ISO 4217 XPF.
var Xpf = New("XPF", "₣", 953).
	WithDecimals(0).
	WithBanknotes(500, 1000, 5000, 10000).
	WithCoins(1, 2, 5, 10, 20, 50, 100)
