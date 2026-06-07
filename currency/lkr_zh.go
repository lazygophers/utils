//go:build country_all || country_asia || country_lk || country_southern_asia || currency_all || currency_lkr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lkr.RegisterName(xlanguage.Chinese, "斯里兰卡卢比")
}
