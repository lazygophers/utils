//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Korean, "카메룬")
	dataCameroon.RegisterOfficialName(xlanguage.Korean, "카메룬 공화국")
	dataCameroon.RegisterCapital(xlanguage.Korean, "야운데")
}
