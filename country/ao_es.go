//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Spanish, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.Spanish, "República de Angola")
	dataAngola.RegisterCapital(xlanguage.Spanish, "Luanda")
}
