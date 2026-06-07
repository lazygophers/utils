//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.French, "Salvador")
	dataElSalvador.RegisterOfficialName(xlanguage.French, "République du Salvador")
	dataElSalvador.RegisterCapital(xlanguage.French, "San Salvador")
}
