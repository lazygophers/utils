//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Spanish, "Singapur")
	dataSingapore.RegisterOfficialName(xlanguage.Spanish, "República de Singapur")
	dataSingapore.RegisterCapital(xlanguage.Spanish, "Singapur")
}
