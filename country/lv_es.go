//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Spanish, "Letonia")
	dataLatvia.RegisterOfficialName(xlanguage.Spanish, "República de Letonia")
	dataLatvia.RegisterCapital(xlanguage.Spanish, "Riga")
}
