//go:build (lang_ko || lang_all) && (country_all || country_europe || country_fi || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Korean, "핀란드")
	dataFinland.RegisterOfficialName(xlanguage.Korean, "핀란드 공화국")
	dataFinland.RegisterCapital(xlanguage.Korean, "헬싱키")
}
