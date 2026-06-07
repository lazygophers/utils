//go:build country_all || country_europe || country_gi || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.English, "Gibraltar")
	dataGibraltar.RegisterOfficialName(xlanguage.English, "Gibraltar")
	dataGibraltar.RegisterCapital(xlanguage.English, "Gibraltar")
}
