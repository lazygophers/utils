//go:build (lang_fr || lang_all) && (country_all || country_asia || country_central_asia || country_uz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.French, "Ouzbékistan")
	dataUzbekistan.RegisterOfficialName(xlanguage.French, "République d'Ouzbékistan")
	dataUzbekistan.RegisterCapital(xlanguage.French, "Tachkent")
}
