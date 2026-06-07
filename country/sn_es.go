//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Spanish, "Senegal")
	dataSenegal.RegisterOfficialName(xlanguage.Spanish, "República de Senegal")
	dataSenegal.RegisterCapital(xlanguage.Spanish, "Dakar")
}
