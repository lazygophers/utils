//go:build (lang_ko || lang_all) && (country_all || country_asia || country_iq || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Korean, "이라크")
	dataIraq.RegisterOfficialName(xlanguage.Korean, "이라크 공화국")
	dataIraq.RegisterCapital(xlanguage.Korean, "바그다드")
}
