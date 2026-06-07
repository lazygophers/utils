//go:build (lang_ko || lang_all) && (country_africa || country_all || country_dz || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Korean, "알제리")
	dataAlgeria.RegisterOfficialName(xlanguage.Korean, "알제리 인민 민주 공화국")
	dataAlgeria.RegisterCapital(xlanguage.Korean, "알제")
}
