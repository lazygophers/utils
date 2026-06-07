//go:build (lang_ko || lang_all) && (country_all || country_micronesia || country_oceania || country_um)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Korean, "미국령 군소 제도")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Korean, "미국령 군소 제도")
}
