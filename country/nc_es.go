//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Spanish, "Nueva Caledonia")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Spanish, "Nueva Caledonia")
	dataNewCaledonia.RegisterCapital(xlanguage.Spanish, "Numea")
}
