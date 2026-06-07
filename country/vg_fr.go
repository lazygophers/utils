//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_vg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.French, "Îles Vierges britanniques")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.French, "Îles Vierges britanniques")
	dataBritishVirginIslands.RegisterCapital(xlanguage.French, "Road Town")
}
