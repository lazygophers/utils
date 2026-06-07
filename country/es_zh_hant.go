//go:build (lang_zh_hant || lang_all) && (country_all || country_es || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.MustParse("zh-Hant"), "西班牙")
	dataSpain.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "西班牙王國")
	dataSpain.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬德里")
}
