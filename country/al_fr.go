//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.French, "Albanie")
	dataAlbania.RegisterOfficialName(xlanguage.French, "République d'Albanie")
	dataAlbania.RegisterCapital(xlanguage.French, "Tirana")
}
