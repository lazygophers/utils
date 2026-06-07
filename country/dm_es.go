//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Spanish, "Dominica")
	dataDominica.RegisterOfficialName(xlanguage.Spanish, "Mancomunidad de Dominica")
	dataDominica.RegisterCapital(xlanguage.Spanish, "Roseau")
}
