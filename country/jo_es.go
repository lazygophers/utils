//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Spanish, "Jordania")
	dataJordan.RegisterOfficialName(xlanguage.Spanish, "Reino Hachemita de Jordania")
	dataJordan.RegisterCapital(xlanguage.Spanish, "Ammán")
}
