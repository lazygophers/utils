//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_tc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Spanish, "Islas Turcas y Caicos")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Turcas y Caicos")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Spanish, "Cockburn Town")
}
