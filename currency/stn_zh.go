//go:build country_africa || country_all || country_middle_africa || country_st || currency_all || currency_stn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	STN.RegisterName(xlanguage.Chinese, "圣多美和普林西比多布拉")
}
