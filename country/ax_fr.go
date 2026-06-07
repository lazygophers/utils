//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.French, "Åland")
	dataAlandIslands.RegisterOfficialName(xlanguage.French, "Åland")
	dataAlandIslands.RegisterCapital(xlanguage.French, "Mariehamn")
}
