//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_vi)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.French, "Îles Vierges des États-Unis")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.French, "Îles Vierges des États-Unis")
	dataUsVirginIslands.RegisterCapital(xlanguage.French, "Charlotte-Amalie")
}
