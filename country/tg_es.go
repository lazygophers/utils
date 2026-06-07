//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Spanish, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.Spanish, "República Togolesa")
	dataTogo.RegisterCapital(xlanguage.Spanish, "Lomé")
}
