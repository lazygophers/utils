//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_gm || country_western_africa || currency_all || currency_gmd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GMD.RegisterName(xlanguage.MustParse("zh-Hant"), "達拉西")
}
