//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Korean, "벨리즈")
	dataBelize.RegisterOfficialName(xlanguage.Korean, "벨리즈")
	dataBelize.RegisterCapital(xlanguage.Korean, "벨모판")
}
