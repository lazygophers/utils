//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.French, "Équateur")
	dataEcuador.RegisterOfficialName(xlanguage.French, "République de l'Équateur")
	dataEcuador.RegisterCapital(xlanguage.French, "Quito")
}
