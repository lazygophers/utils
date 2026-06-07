//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Spanish, "Albania")
	dataAlbania.RegisterOfficialName(xlanguage.Spanish, "República de Albania")
	dataAlbania.RegisterCapital(xlanguage.Spanish, "Tirana")
}
