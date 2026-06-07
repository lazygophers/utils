//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Spanish, "Zambia")
	dataZambia.RegisterOfficialName(xlanguage.Spanish, "República de Zambia")
	dataZambia.RegisterCapital(xlanguage.Spanish, "Lusaka")
}
