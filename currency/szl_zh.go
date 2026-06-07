//go:build country_africa || country_all || country_southern_africa || country_sz || currency_all || currency_szl

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SZL.RegisterName(xlanguage.Chinese, "斯威士兰里兰吉尼")
}
