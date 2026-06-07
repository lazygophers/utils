//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Korean, "칠레")
	dataChile.RegisterOfficialName(xlanguage.Korean, "칠레 공화국")
	dataChile.RegisterCapital(xlanguage.Korean, "산티아고")
}
