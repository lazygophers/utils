//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_kw || country_western_asia || currency_all || currency_kwd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KWD.RegisterName(xlanguage.MustParse("zh-Hant"), "科威特第納爾")
}
