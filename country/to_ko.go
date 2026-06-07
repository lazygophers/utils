//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Korean, "통가")
	dataTonga.RegisterOfficialName(xlanguage.Korean, "통가 왕국")
	dataTonga.RegisterCapital(xlanguage.Korean, "누쿠알로파")
}
