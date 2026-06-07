//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_sa || country_western_asia || currency_all || currency_sar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sar.RegisterName(xlanguage.MustParse("zh-Hant"), "沙烏地里亞爾")
}
