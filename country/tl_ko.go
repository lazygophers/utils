//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Korean, "동티모르")
	dataTimorLeste.RegisterOfficialName(xlanguage.Korean, "동티모르 민주 공화국")
	dataTimorLeste.RegisterCapital(xlanguage.Korean, "딜리")
}
