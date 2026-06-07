//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Spanish, "Tanzania")
	dataTanzania.RegisterOfficialName(xlanguage.Spanish, "República Unida de Tanzania")
	dataTanzania.RegisterCapital(xlanguage.Spanish, "Dodoma")
}
