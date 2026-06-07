//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Korean, "그리스")
	dataGreece.RegisterOfficialName(xlanguage.Korean, "그리스 공화국")
	dataGreece.RegisterCapital(xlanguage.Korean, "아테네")
}
