//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Korean, "루마니아")
	dataRomania.RegisterOfficialName(xlanguage.Korean, "루마니아")
	dataRomania.RegisterCapital(xlanguage.Korean, "부쿠레슈티")
}
