//go:build (lang_es || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Spanish, "Corea del Norte")
	dataNorthKorea.RegisterOfficialName(xlanguage.Spanish, "República Popular Democrática de Corea")
	dataNorthKorea.RegisterCapital(xlanguage.Spanish, "Pionyang")
}
