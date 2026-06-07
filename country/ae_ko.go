//go:build (lang_ko || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Korean, "아랍에미리트")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Korean, "아랍에미리트")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Korean, "아부다비")
}
