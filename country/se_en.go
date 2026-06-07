//go:build country_all || country_europe || country_northern_europe || country_se

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.English, "Sweden")
	dataSweden.RegisterOfficialName(xlanguage.English, "Kingdom of Sweden")
	dataSweden.RegisterCapital(xlanguage.English, "Stockholm")
}
