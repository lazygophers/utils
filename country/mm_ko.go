//go:build (lang_ko || lang_all) && (country_all || country_asia || country_mm || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Korean, "미얀마")
	dataMyanmar.RegisterOfficialName(xlanguage.Korean, "미얀마 연방 공화국")
	dataMyanmar.RegisterCapital(xlanguage.Korean, "네피도")
}
