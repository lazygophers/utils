//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Spanish, "San Bartolomé")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Spanish, "Colectividad de San Bartolomé")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Spanish, "Gustavia")
}
