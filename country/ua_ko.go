//go:build (lang_ko || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Korean, "우크라이나")
	dataUkraine.RegisterOfficialName(xlanguage.Korean, "우크라이나")
	dataUkraine.RegisterCapital(xlanguage.Korean, "키이우")
}
