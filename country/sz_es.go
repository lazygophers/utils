//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Spanish, "Esuatini")
	dataEswatini.RegisterOfficialName(xlanguage.Spanish, "Reino de Esuatini")
	dataEswatini.RegisterCapital(xlanguage.Spanish, "Mbabane")
}
