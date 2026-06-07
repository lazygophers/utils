//go:build country_africa || country_all || country_gh || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.English, "Ghana")
	dataGhana.RegisterOfficialName(xlanguage.English, "Republic of Ghana")
	dataGhana.RegisterCapital(xlanguage.English, "Accra")
}
