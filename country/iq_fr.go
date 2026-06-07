//go:build (lang_fr || lang_all) && (country_all || country_asia || country_iq || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.French, "Irak")
	dataIraq.RegisterOfficialName(xlanguage.French, "République d'Irak")
	dataIraq.RegisterCapital(xlanguage.French, "Bagdad")
}
