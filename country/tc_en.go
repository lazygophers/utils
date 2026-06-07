//go:build country_all || country_americas || country_caribbean || country_tc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.English, "Turks and Caicos Islands")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.English, "Turks and Caicos Islands")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.English, "Cockburn Town")
}
