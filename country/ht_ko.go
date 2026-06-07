//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_ht)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Korean, "아이티")
	dataHaiti.RegisterOfficialName(xlanguage.Korean, "아이티 공화국")
	dataHaiti.RegisterCapital(xlanguage.Korean, "포르토프랭스")
}
