//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Korean, "남수단")
	dataSouthSudan.RegisterOfficialName(xlanguage.Korean, "남수단 공화국")
	dataSouthSudan.RegisterCapital(xlanguage.Korean, "주바")
}
