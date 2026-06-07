//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Spanish, "Sudáfrica")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Spanish, "República de Sudáfrica")
	dataSouthAfrica.RegisterCapital(xlanguage.Spanish, "Pretoria")
}
