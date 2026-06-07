//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Spanish, "Sudán del Sur")
	dataSouthSudan.RegisterOfficialName(xlanguage.Spanish, "República de Sudán del Sur")
	dataSouthSudan.RegisterCapital(xlanguage.Spanish, "Yuba")
}
