//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ao || country_middle_africa || currency_all || currency_aoa)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AOA.RegisterName(xlanguage.MustParse("zh-Hant"), "寬扎")
}
