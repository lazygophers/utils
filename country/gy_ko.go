//go:build (lang_ko || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Korean, "가이아나")
	dataGuyana.RegisterOfficialName(xlanguage.Korean, "가이아나 협동 공화국")
	dataGuyana.RegisterCapital(xlanguage.Korean, "조지타운")
}
