//go:build (lang_fr || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.French, "Émirats arabes unis")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.French, "Émirats arabes unis")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.French, "Abou Dabi")
}
