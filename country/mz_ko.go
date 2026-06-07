//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Korean, "모잠비크")
	dataMozambique.RegisterOfficialName(xlanguage.Korean, "모잠비크 공화국")
	dataMozambique.RegisterCapital(xlanguage.Korean, "마푸투")
}
