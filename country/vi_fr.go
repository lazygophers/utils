//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.French, "Îles Vierges des États-Unis")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.French, "Îles Vierges des États-Unis")
	dataUsVirginIslands.RegisterCapital(xlanguage.French, "Charlotte-Amalie")
}
