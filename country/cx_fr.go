//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.French, "Île Christmas")
	dataChristmasIsland.RegisterOfficialName(xlanguage.French, "Territoire de l'île Christmas")
	dataChristmasIsland.RegisterCapital(xlanguage.French, "Flying Fish Cove")
}
