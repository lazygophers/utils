//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bs || country_caribbean || currency_all || currency_bsd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BSD.RegisterName(xlanguage.MustParse("zh-Hant"), "巴哈馬元")
}
