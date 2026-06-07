//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_ky || currency_all || currency_kyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KYD.RegisterName(xlanguage.MustParse("zh-Hant"), "開曼群島元")
}
