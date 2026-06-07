//go:build country_africa || country_all || country_ng || country_western_africa || currency_all || currency_ngn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ngn.RegisterName(xlanguage.Chinese, "尼日利亚奈拉")
}
