//go:build (lang_fr || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_tl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.French, "Timor oriental")
	dataTimorLeste.RegisterOfficialName(xlanguage.French, "République démocratique du Timor oriental")
	dataTimorLeste.RegisterCapital(xlanguage.French, "Dili")
}
