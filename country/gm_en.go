//go:build country_africa || country_all || country_gm || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.English, "Gambia")
	dataGambia.RegisterOfficialName(xlanguage.English, "Republic of the Gambia")
	dataGambia.RegisterCapital(xlanguage.English, "Banjul")
}
