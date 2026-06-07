//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Spanish, "Granada")
	dataGrenada.RegisterOfficialName(xlanguage.Spanish, "Granada")
	dataGrenada.RegisterCapital(xlanguage.Spanish, "Saint George's")
}
