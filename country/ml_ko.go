//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Korean, "말리")
	dataMali.RegisterOfficialName(xlanguage.Korean, "말리 공화국")
	dataMali.RegisterCapital(xlanguage.Korean, "바마코")
}
