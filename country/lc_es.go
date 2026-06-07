//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Spanish, "Santa Lucía")
	dataSaintLucia.RegisterOfficialName(xlanguage.Spanish, "Santa Lucía")
	dataSaintLucia.RegisterCapital(xlanguage.Spanish, "Castries")
}
