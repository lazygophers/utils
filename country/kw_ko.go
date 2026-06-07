//go:build (lang_ko || lang_all) && (country_all || country_asia || country_kw || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Korean, "쿠웨이트")
	dataKuwait.RegisterOfficialName(xlanguage.Korean, "쿠웨이트국")
	dataKuwait.RegisterCapital(xlanguage.Korean, "쿠웨이트시티")
}
