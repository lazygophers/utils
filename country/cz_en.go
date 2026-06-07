//go:build country_all || country_cz || country_eastern_europe || country_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.English, "Czechia")
	dataCzechia.RegisterOfficialName(xlanguage.English, "Czech Republic")
	dataCzechia.RegisterCapital(xlanguage.English, "Prague")
}
