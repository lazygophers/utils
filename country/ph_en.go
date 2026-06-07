//go:build country_all || country_asia || country_ph || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.English, "Philippines")
	dataPhilippines.RegisterOfficialName(xlanguage.English, "Republic of the Philippines")
	dataPhilippines.RegisterCapital(xlanguage.English, "Manila")
}
