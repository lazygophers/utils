//go:build (lang_ko || lang_all) && (country_all || country_at || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Korean, "오스트리아")
	dataAustria.RegisterOfficialName(xlanguage.Korean, "오스트리아 공화국")
	dataAustria.RegisterCapital(xlanguage.Korean, "빈")
}
