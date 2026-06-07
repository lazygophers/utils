//go:build (lang_fr || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.French, "Philippines")
	dataPhilippines.RegisterOfficialName(xlanguage.French, "République des Philippines")
	dataPhilippines.RegisterCapital(xlanguage.French, "Manille")
}
