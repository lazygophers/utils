//go:build country_all || country_asia || country_western_asia || country_ye || currency_all || currency_yer

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	YER.RegisterName(xlanguage.Chinese, "也门里亚尔")
}
