//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_dz || country_northern_africa || currency_all || currency_dzd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dzd.RegisterName(xlanguage.MustParse("zh-Hant"), "阿爾及利亞第納爾")
}
