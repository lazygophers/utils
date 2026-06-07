//go:build (lang_ko || lang_all) && (country_all || country_europe || country_southern_europe || country_va)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Korean, "바티칸 시국")
	dataVaticanCity.RegisterOfficialName(xlanguage.Korean, "바티칸 시국")
	dataVaticanCity.RegisterCapital(xlanguage.Korean, "바티칸시")
}
