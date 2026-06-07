//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Korean, "에콰도르")
	dataEcuador.RegisterOfficialName(xlanguage.Korean, "에콰도르 공화국")
	dataEcuador.RegisterCapital(xlanguage.Korean, "키토")
}
