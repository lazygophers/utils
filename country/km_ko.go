//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_km)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Korean, "코모로")
	dataComoros.RegisterOfficialName(xlanguage.Korean, "코모로 연합")
	dataComoros.RegisterCapital(xlanguage.Korean, "모로니")
}
