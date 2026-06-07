//go:build (lang_es || lang_all) && (country_all || country_asia || country_mv || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Spanish, "Maldivas")
	dataMaldives.RegisterOfficialName(xlanguage.Spanish, "República de Maldivas")
	dataMaldives.RegisterCapital(xlanguage.Spanish, "Malé")
}
