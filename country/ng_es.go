//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Spanish, "Nigeria")
	dataNigeria.RegisterOfficialName(xlanguage.Spanish, "República Federal de Nigeria")
	dataNigeria.RegisterCapital(xlanguage.Spanish, "Abuya")
}
