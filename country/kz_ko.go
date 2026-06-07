//go:build (lang_ko || lang_all) && (country_all || country_asia || country_central_asia || country_kz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Korean, "카자흐스탄")
	dataKazakhstan.RegisterOfficialName(xlanguage.Korean, "카자흐스탄 공화국")
	dataKazakhstan.RegisterCapital(xlanguage.Korean, "아스타나")
}
