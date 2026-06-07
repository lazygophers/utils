//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Spanish, "Mauricio")
	dataMauritius.RegisterOfficialName(xlanguage.Spanish, "República de Mauricio")
	dataMauritius.RegisterCapital(xlanguage.Spanish, "Port Louis")
}
