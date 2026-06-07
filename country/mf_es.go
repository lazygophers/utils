//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Spanish, "San Martín")
	dataSaintMartin.RegisterOfficialName(xlanguage.Spanish, "Colectividad de San Martín")
	dataSaintMartin.RegisterCapital(xlanguage.Spanish, "Marigot")
}
