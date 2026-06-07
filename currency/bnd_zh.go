//go:build country_all || country_asia || country_bn || country_south_eastern_asia || currency_all || currency_bnd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BND.RegisterName(xlanguage.Chinese, "文莱元")
}
