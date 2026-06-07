//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Korean, "차드")
	dataChad.RegisterOfficialName(xlanguage.Korean, "차드 공화국")
	dataChad.RegisterCapital(xlanguage.Korean, "은자메나")
}
