//go:build country_africa || country_all || country_na || country_southern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.English, "Namibia")
	dataNamibia.RegisterOfficialName(xlanguage.English, "Republic of Namibia")
	dataNamibia.RegisterCapital(xlanguage.English, "Windhoek")
}
