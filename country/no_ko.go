//go:build (lang_ko || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Korean, "노르웨이")
	dataNorway.RegisterOfficialName(xlanguage.Korean, "노르웨이 왕국")
	dataNorway.RegisterCapital(xlanguage.Korean, "오슬로")
}
