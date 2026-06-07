//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Korean, "피지")
	dataFiji.RegisterOfficialName(xlanguage.Korean, "피지 공화국")
	dataFiji.RegisterCapital(xlanguage.Korean, "수바")
}
