//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Spanish, "República Democrática del Congo")
	dataDrCongo.RegisterOfficialName(xlanguage.Spanish, "República Democrática del Congo")
	dataDrCongo.RegisterCapital(xlanguage.Spanish, "Kinsasa")
}
