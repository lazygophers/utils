//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Spanish, "Marruecos")
	dataMorocco.RegisterOfficialName(xlanguage.Spanish, "Reino de Marruecos")
	dataMorocco.RegisterCapital(xlanguage.Spanish, "Rabat")
}
