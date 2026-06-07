//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Spanish, "Islas Åland")
	dataAlandIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Åland")
	dataAlandIslands.RegisterCapital(xlanguage.Spanish, "Mariehamn")
}
