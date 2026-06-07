//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_na || country_southern_africa || currency_all || currency_nad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NAD.RegisterName(xlanguage.MustParse("zh-Hant"), "納米比亞元")
}
