//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ly || country_northern_africa || currency_all || currency_lyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LYD.RegisterName(xlanguage.MustParse("zh-Hant"), "利比亞第納爾")
}
