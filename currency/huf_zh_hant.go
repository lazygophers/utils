//go:build (lang_zh_hant || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu || currency_all || currency_huf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	HUF.RegisterName(xlanguage.MustParse("zh-Hant"), "福林")
}
