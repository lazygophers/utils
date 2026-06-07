//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Korean, "아랍에미리트")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Korean, "아랍에미리트")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Korean, "아부다비")
}
