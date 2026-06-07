//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.Spanish, "Botsuana")
	dataBotswana.RegisterOfficialName(xlanguage.Spanish, "República de Botsuana")
	dataBotswana.RegisterCapital(xlanguage.Spanish, "Gaborone")
}
