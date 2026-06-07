//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Korean, "케이맨 제도")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Korean, "케이맨 제도")
	dataCaymanIslands.RegisterCapital(xlanguage.Korean, "조지타운")
}
