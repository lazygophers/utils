//go:build (lang_ko || lang_all) && (country_all || country_asia || country_sy || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Korean, "시리아")
	dataSyria.RegisterOfficialName(xlanguage.Korean, "시리아 아랍 공화국")
	dataSyria.RegisterCapital(xlanguage.Korean, "다마스쿠스")
}
