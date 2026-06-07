//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.French, "Bolivie")
	dataBolivia.RegisterOfficialName(xlanguage.French, "État plurinational de Bolivie")
	dataBolivia.RegisterCapital(xlanguage.French, "Sucre")
}
