//go:build country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MZN.RegisterName(xlanguage.Chinese, "莫桑比克梅蒂卡尔")
}
