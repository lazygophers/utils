//go:build country_all || country_asia || country_mm || country_south_eastern_asia || currency_all || currency_mmk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MMK.RegisterName(xlanguage.Chinese, "缅甸元")
}
