//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_pk || country_southern_asia || currency_all || currency_pkr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pkr.RegisterName(xlanguage.MustParse("zh-Hant"), "巴基斯坦盧比")
}
