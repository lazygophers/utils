//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Korean, "버뮤다")
	dataBermuda.RegisterOfficialName(xlanguage.Korean, "버뮤다")
	dataBermuda.RegisterCapital(xlanguage.Korean, "해밀턴")
}
