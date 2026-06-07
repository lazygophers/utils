//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_km || currency_all || currency_kmf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kmf.RegisterName(xlanguage.MustParse("zh-Hant"), "葛摩法郎")
}
