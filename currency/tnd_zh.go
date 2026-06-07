//go:build country_africa || country_all || country_northern_africa || country_tn || currency_all || currency_tnd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TND.RegisterName(xlanguage.Chinese, "突尼斯第纳尔")
}
