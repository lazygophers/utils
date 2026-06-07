//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_om || country_western_asia || currency_all || currency_omr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Omr.RegisterName(xlanguage.MustParse("zh-Hant"), "阿曼里亞爾")
}
