//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_mv || country_southern_asia || currency_all || currency_mvr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MVR.RegisterName(xlanguage.MustParse("zh-Hant"), "馬爾地夫拉菲亞")
}
