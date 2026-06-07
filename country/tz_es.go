//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Spanish, "Tanzania")
	dataTanzania.RegisterOfficialName(xlanguage.Spanish, "República Unida de Tanzania")
	dataTanzania.RegisterCapital(xlanguage.Spanish, "Dodoma")
}
