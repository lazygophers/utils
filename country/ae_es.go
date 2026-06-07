//go:build (lang_es || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Spanish, "Emiratos Árabes Unidos")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Spanish, "Emiratos Árabes Unidos")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Spanish, "Abu Dabi")
}
