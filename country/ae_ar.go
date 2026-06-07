//go:build country_ae || country_all || country_asia || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Arabic, "الإمارات العربية المتحدة")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Arabic, "الإمارات العربية المتحدة")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Arabic, "أبوظبي")
}
