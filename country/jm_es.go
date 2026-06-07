//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Spanish, "Jamaica")
	dataJamaica.RegisterOfficialName(xlanguage.Spanish, "Jamaica")
	dataJamaica.RegisterCapital(xlanguage.Spanish, "Kingston")
}
