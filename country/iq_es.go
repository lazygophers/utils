//go:build (lang_es || lang_all) && (country_all || country_asia || country_iq || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Spanish, "Irak")
	dataIraq.RegisterOfficialName(xlanguage.Spanish, "República de Irak")
	dataIraq.RegisterCapital(xlanguage.Spanish, "Bagdad")
}
