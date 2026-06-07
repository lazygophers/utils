//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Spanish, "Malí")
	dataMali.RegisterOfficialName(xlanguage.Spanish, "República de Malí")
	dataMali.RegisterCapital(xlanguage.Spanish, "Bamako")
}
