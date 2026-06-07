//go:build (lang_ko || lang_all) && (country_all || country_asia || country_central_asia || country_tj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Korean, "타지키스탄")
	dataTajikistan.RegisterOfficialName(xlanguage.Korean, "타지키스탄 공화국")
	dataTajikistan.RegisterCapital(xlanguage.Korean, "두샨베")
}
