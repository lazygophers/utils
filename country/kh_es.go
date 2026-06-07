//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Spanish, "Camboya")
	dataCambodia.RegisterOfficialName(xlanguage.Spanish, "Reino de Camboya")
	dataCambodia.RegisterCapital(xlanguage.Spanish, "Nom Pen")
}
