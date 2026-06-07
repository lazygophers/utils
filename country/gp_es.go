//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Spanish, "Guadalupe")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Spanish, "Guadalupe")
	dataGuadeloupe.RegisterCapital(xlanguage.Spanish, "Basse-Terre")
}
