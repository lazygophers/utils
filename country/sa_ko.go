//go:build (lang_ko || lang_all) && (country_all || country_asia || country_sa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Korean, "사우디아라비아")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Korean, "사우디아라비아 왕국")
	dataSaudiArabia.RegisterCapital(xlanguage.Korean, "리야드")
}
