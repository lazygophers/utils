//go:build country_all || country_eastern_europe || country_europe || country_ua

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.English, "Ukraine")
	dataUkraine.RegisterOfficialName(xlanguage.English, "Ukraine")
	dataUkraine.RegisterCapital(xlanguage.English, "Kyiv")
}
