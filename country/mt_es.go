//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.Spanish, "Malta")
	dataMalta.RegisterOfficialName(xlanguage.Spanish, "República de Malta")
	dataMalta.RegisterCapital(xlanguage.Spanish, "La Valeta")
}
