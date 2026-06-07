//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_dm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Spanish, "Dominica")
	dataDominica.RegisterOfficialName(xlanguage.Spanish, "Mancomunidad de Dominica")
	dataDominica.RegisterCapital(xlanguage.Spanish, "Roseau")
}
