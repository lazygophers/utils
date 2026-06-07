//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_tc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.French, "Îles Turques-et-Caïques")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.French, "Îles Turques-et-Caïques")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.French, "Cockburn Town")
}
