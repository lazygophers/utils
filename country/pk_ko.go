//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Korean, "파키스탄")
	dataPakistan.RegisterOfficialName(xlanguage.Korean, "파키스탄 이슬람 공화국")
	dataPakistan.RegisterCapital(xlanguage.Korean, "이슬라마바드")
}
