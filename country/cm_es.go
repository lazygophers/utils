//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Spanish, "Camerún")
	dataCameroon.RegisterOfficialName(xlanguage.Spanish, "República de Camerún")
	dataCameroon.RegisterCapital(xlanguage.Spanish, "Yaundé")
}
