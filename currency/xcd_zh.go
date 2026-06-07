//go:build country_ag || country_ai || country_all || country_americas || country_caribbean || country_dm || country_gd || country_kn || country_lc || country_ms || country_vc || currency_all || currency_xcd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	XCD.RegisterName(xlanguage.Chinese, "东加勒比元")
}
