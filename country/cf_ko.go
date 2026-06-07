//go:build (lang_ko || lang_all) && (country_africa || country_all || country_cf || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Korean, "중앙아프리카 공화국")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Korean, "중앙아프리카 공화국")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Korean, "방기")
}
