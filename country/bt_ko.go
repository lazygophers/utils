//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Korean, "부탄")
	dataBhutan.RegisterOfficialName(xlanguage.Korean, "부탄 왕국")
	dataBhutan.RegisterCapital(xlanguage.Korean, "팀부")
}
