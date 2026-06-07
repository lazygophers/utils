//go:build country_africa || country_all || country_eastern_africa || country_tz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.English, "Tanzania")
	dataTanzania.RegisterOfficialName(xlanguage.English, "United Republic of Tanzania")
	dataTanzania.RegisterCapital(xlanguage.English, "Dodoma")
}
