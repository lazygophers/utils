//go:build country_all || country_cz || country_eastern_europe || country_europe || currency_all || currency_czk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CZK.RegisterName(xlanguage.Chinese, "捷克克朗")
}
