//go:build country_ae || country_all || country_asia || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.English, "United Arab Emirates")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.English, "United Arab Emirates")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.English, "Abu Dhabi")
}
