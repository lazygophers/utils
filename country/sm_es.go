//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Spanish, "San Marino")
	dataSanMarino.RegisterOfficialName(xlanguage.Spanish, "República de San Marino")
	dataSanMarino.RegisterCapital(xlanguage.Spanish, "San Marino")
}
