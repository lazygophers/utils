//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Spanish, "Malaui")
	dataMalawi.RegisterOfficialName(xlanguage.Spanish, "República de Malaui")
	dataMalawi.RegisterCapital(xlanguage.Spanish, "Lilongüe")
}
