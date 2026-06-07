//go:build country_all || country_asia || country_ir || country_southern_asia || currency_all || currency_irr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Irr.RegisterName(xlanguage.Chinese, "伊朗里亚尔")
}
