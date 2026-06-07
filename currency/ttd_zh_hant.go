//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ttd.RegisterName(xlanguage.MustParse("zh-Hant"), "千里達及托巴哥元")
}
