//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Korean, "튀니지")
	dataTunisia.RegisterOfficialName(xlanguage.Korean, "튀니지 공화국")
	dataTunisia.RegisterCapital(xlanguage.Korean, "튀니스")
}
