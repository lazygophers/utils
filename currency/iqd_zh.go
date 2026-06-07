//go:build country_all || country_asia || country_iq || country_western_asia || currency_all || currency_iqd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	IQD.RegisterName(xlanguage.Chinese, "伊拉克第纳尔")
}
