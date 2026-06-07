//go:build (lang_ko || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Korean, "슬로베니아")
	dataSlovenia.RegisterOfficialName(xlanguage.Korean, "슬로베니아 공화국")
	dataSlovenia.RegisterCapital(xlanguage.Korean, "류블랴나")
}
