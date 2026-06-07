//go:build country_all || country_asia || country_la || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.English, "Laos")
	dataLaos.RegisterOfficialName(xlanguage.English, "Lao People's Democratic Republic")
	dataLaos.RegisterCapital(xlanguage.English, "Vientiane")
}
