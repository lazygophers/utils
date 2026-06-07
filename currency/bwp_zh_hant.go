//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_bw || country_southern_africa || currency_all || currency_bwp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bwp.RegisterName(xlanguage.MustParse("zh-Hant"), "普拉")
}
