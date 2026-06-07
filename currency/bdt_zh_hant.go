//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bd || country_southern_asia || currency_all || currency_bdt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BDT.RegisterName(xlanguage.MustParse("zh-Hant"), "塔卡")
}
