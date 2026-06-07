//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc || currency_all || currency_scr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Scr.RegisterName(xlanguage.MustParse("zh-Hant"), "塞席爾盧比")
}
