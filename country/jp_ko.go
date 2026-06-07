//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.Korean, "일본")
	dataJapan.RegisterOfficialName(xlanguage.Korean, "일본국")
	dataJapan.RegisterCapital(xlanguage.Korean, "도쿄")
}
