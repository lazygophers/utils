//go:build (lang_es || lang_all) && (country_all || country_asia || country_bd || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.Spanish, "Bangladés")
	dataBangladesh.RegisterOfficialName(xlanguage.Spanish, "República Popular de Bangladés")
	dataBangladesh.RegisterCapital(xlanguage.Spanish, "Daca")
}
