//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Spanish, "Nueva Zelanda")
	dataNewZealand.RegisterOfficialName(xlanguage.Spanish, "Nueva Zelanda")
	dataNewZealand.RegisterCapital(xlanguage.Spanish, "Wellington")
}
