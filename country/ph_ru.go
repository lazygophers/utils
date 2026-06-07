//go:build (lang_ru || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Russian, "Филиппины")
	dataPhilippines.RegisterOfficialName(xlanguage.Russian, "Республика Филиппины")
	dataPhilippines.RegisterCapital(xlanguage.Russian, "Манила")
}
