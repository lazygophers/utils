//go:build (lang_ko || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Korean, "라트비아")
	dataLatvia.RegisterOfficialName(xlanguage.Korean, "라트비아 공화국")
	dataLatvia.RegisterCapital(xlanguage.Korean, "리가")
}
