//go:build (lang_ar || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Arabic, "الفلبين")
	dataPhilippines.RegisterOfficialName(xlanguage.Arabic, "جمهورية الفلبين")
	dataPhilippines.RegisterCapital(xlanguage.Arabic, "مانيلا")
}
