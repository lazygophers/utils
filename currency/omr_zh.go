//go:build country_all || country_asia || country_om || country_western_asia || currency_all || currency_omr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	OMR.RegisterName(xlanguage.Chinese, "阿曼里亚尔")
}
