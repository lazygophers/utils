//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Korean, "마다가스카르")
	dataMadagascar.RegisterOfficialName(xlanguage.Korean, "마다가스카르 공화국")
	dataMadagascar.RegisterCapital(xlanguage.Korean, "안타나나리보")
}
