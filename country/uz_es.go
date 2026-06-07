//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Spanish, "Uzbekistán")
	dataUzbekistan.RegisterOfficialName(xlanguage.Spanish, "República de Uzbekistán")
	dataUzbekistan.RegisterCapital(xlanguage.Spanish, "Taskent")
}
