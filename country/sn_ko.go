//go:build (lang_ko || lang_all) && (country_africa || country_all || country_sn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Korean, "세네갈")
	dataSenegal.RegisterOfficialName(xlanguage.Korean, "세네갈 공화국")
	dataSenegal.RegisterCapital(xlanguage.Korean, "다카르")
}
