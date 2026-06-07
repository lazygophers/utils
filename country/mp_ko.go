//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Korean, "북마리아나 제도")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Korean, "북마리아나 제도 연방")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Korean, "사이판")
}
