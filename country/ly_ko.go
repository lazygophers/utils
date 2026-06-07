//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ly || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.Korean, "리비아")
	dataLibya.RegisterOfficialName(xlanguage.Korean, "리비아국")
	dataLibya.RegisterCapital(xlanguage.Korean, "트리폴리")
}
