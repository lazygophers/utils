//go:build country_all || country_asia || country_south_eastern_asia || country_th

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.English, "Thailand")
	dataThailand.RegisterOfficialName(xlanguage.English, "Kingdom of Thailand")
	dataThailand.RegisterCapital(xlanguage.English, "Bangkok")
}
