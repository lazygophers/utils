//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia || currency_all || currency_khr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Khr.RegisterName(xlanguage.MustParse("zh-Hant"), "瑞爾")
}
