//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.French, "Îles Malouines")
	dataFalklandIslands.RegisterOfficialName(xlanguage.French, "Îles Malouines")
	dataFalklandIslands.RegisterCapital(xlanguage.French, "Stanley")
}
