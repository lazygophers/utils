//go:build (lang_ko || lang_all) && (country_all || country_americas || country_ar || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Korean, "아르헨티나")
	dataArgentina.RegisterOfficialName(xlanguage.Korean, "아르헨티나 공화국")
	dataArgentina.RegisterCapital(xlanguage.Korean, "부에노스아이레스")
}
