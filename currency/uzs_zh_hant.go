//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_central_asia || country_uz || currency_all || currency_uzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	UZS.RegisterName(xlanguage.MustParse("zh-Hant"), "烏茲別克索姆")
}
