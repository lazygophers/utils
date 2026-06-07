//go:build (lang_ko || lang_all) && (country_africa || country_all || country_northern_africa || country_sd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Korean, "수단")
	dataSudan.RegisterOfficialName(xlanguage.Korean, "수단 공화국")
	dataSudan.RegisterCapital(xlanguage.Korean, "하르툼")
}
