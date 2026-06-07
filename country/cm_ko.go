//go:build (lang_ko || lang_all) && (country_africa || country_all || country_cm || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Korean, "카메룬")
	dataCameroon.RegisterOfficialName(xlanguage.Korean, "카메룬 공화국")
	dataCameroon.RegisterCapital(xlanguage.Korean, "야운데")
}
