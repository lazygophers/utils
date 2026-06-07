//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Korean, "요르단")
	dataJordan.RegisterOfficialName(xlanguage.Korean, "요르단 하심 왕국")
	dataJordan.RegisterCapital(xlanguage.Korean, "암만")
}
