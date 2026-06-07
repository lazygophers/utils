//go:build (lang_ko || lang_all) && (country_africa || country_all || country_gh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Korean, "가나")
	dataGhana.RegisterOfficialName(xlanguage.Korean, "가나 공화국")
	dataGhana.RegisterCapital(xlanguage.Korean, "아크라")
}
