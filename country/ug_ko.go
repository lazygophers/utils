//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Korean, "우간다")
	dataUganda.RegisterOfficialName(xlanguage.Korean, "우간다 공화국")
	dataUganda.RegisterCapital(xlanguage.Korean, "캄팔라")
}
