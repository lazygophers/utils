//go:build (lang_ko || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Korean, "키프로스")
	dataCyprus.RegisterOfficialName(xlanguage.Korean, "키프로스 공화국")
	dataCyprus.RegisterCapital(xlanguage.Korean, "니코시아")
}
