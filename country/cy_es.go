//go:build (lang_es || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Spanish, "Chipre")
	dataCyprus.RegisterOfficialName(xlanguage.Spanish, "República de Chipre")
	dataCyprus.RegisterCapital(xlanguage.Spanish, "Nicosia")
}
