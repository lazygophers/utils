//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bt || country_southern_asia || currency_all || currency_btn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BTN.RegisterName(xlanguage.MustParse("zh-Hant"), "不丹努爾特魯姆")
}
