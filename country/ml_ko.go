//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ml || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Korean, "말리")
	dataMali.RegisterOfficialName(xlanguage.Korean, "말리 공화국")
	dataMali.RegisterCapital(xlanguage.Korean, "바마코")
}
