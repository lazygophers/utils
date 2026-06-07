//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_aw || country_caribbean || currency_all || currency_awg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Awg.RegisterName(xlanguage.MustParse("zh-Hant"), "阿魯巴弗羅林")
}
