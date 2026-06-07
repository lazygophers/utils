//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Spanish, "Gibraltar")
	dataGibraltar.RegisterOfficialName(xlanguage.Spanish, "Gibraltar")
	dataGibraltar.RegisterCapital(xlanguage.Spanish, "Gibraltar")
}
