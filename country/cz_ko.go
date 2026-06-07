//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Korean, "체코")
	dataCzechia.RegisterOfficialName(xlanguage.Korean, "체코 공화국")
	dataCzechia.RegisterCapital(xlanguage.Korean, "프라하")
}
