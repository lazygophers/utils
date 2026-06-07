//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_polynesia || country_to || currency_all || currency_top)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TOP.RegisterName(xlanguage.MustParse("zh-Hant"), "潘加")
}
