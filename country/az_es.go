//go:build (lang_es || lang_all) && (country_all || country_asia || country_az || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Spanish, "Azerbaiyán")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Spanish, "República de Azerbaiyán")
	dataAzerbaijan.RegisterCapital(xlanguage.Spanish, "Bakú")
}
