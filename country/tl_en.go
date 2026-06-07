//go:build country_all || country_asia || country_south_eastern_asia || country_tl

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.English, "Timor-Leste")
	dataTimorLeste.RegisterOfficialName(xlanguage.English, "Democratic Republic of Timor-Leste")
	dataTimorLeste.RegisterCapital(xlanguage.English, "Dili")
}
