//go:build (lang_ko || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.Korean, "부룬디")
	dataBurundi.RegisterOfficialName(xlanguage.Korean, "부룬디 공화국")
	dataBurundi.RegisterCapital(xlanguage.Korean, "기테가")
}
