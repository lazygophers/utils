//go:build (lang_ko || lang_all) && (country_all || country_asia || country_lb || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Korean, "레바논")
	dataLebanon.RegisterOfficialName(xlanguage.Korean, "레바논 공화국")
	dataLebanon.RegisterCapital(xlanguage.Korean, "베이루트")
}
