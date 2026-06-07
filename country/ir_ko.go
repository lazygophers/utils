//go:build (lang_ko || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Korean, "이란")
	dataIran.RegisterOfficialName(xlanguage.Korean, "이란 이슬람 공화국")
	dataIran.RegisterCapital(xlanguage.Korean, "테헤란")
}
