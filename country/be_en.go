//go:build country_all || country_be || country_europe || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.English, "Belgium")
	dataBelgium.RegisterOfficialName(xlanguage.English, "Kingdom of Belgium")
	dataBelgium.RegisterCapital(xlanguage.English, "Brussels")
}
