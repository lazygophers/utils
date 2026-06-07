//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Spanish, "Montserrat")
	dataMontserrat.RegisterOfficialName(xlanguage.Spanish, "Montserrat")
	dataMontserrat.RegisterCapital(xlanguage.Spanish, "Brades")
}
