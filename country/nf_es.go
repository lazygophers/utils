//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorfolkIsland.RegisterName(xlanguage.Spanish, "Isla Norfolk")
	dataNorfolkIsland.RegisterOfficialName(xlanguage.Spanish, "Territorio de la Isla Norfolk")
	dataNorfolkIsland.RegisterCapital(xlanguage.Spanish, "Kingston")
}
