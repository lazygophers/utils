//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Russian, "Украина")
	dataUkraine.RegisterOfficialName(xlanguage.Russian, "Украина")
	dataUkraine.RegisterCapital(xlanguage.Russian, "Киев")
}
