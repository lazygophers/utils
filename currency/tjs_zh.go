//go:build country_all || country_asia || country_central_asia || country_tj || currency_all || currency_tjs

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TJS.RegisterName(xlanguage.Chinese, "塔吉克斯坦索莫尼")
}
