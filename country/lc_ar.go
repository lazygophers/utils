//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Arabic, "سانت لوسيا")
	dataSaintLucia.RegisterOfficialName(xlanguage.Arabic, "سانت لوسيا")
	dataSaintLucia.RegisterCapital(xlanguage.Arabic, "كاستريس")
}
