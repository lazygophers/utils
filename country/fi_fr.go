//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.French, "Finlande")
	dataFinland.RegisterOfficialName(xlanguage.French, "République de Finlande")
	dataFinland.RegisterCapital(xlanguage.French, "Helsinki")
}
