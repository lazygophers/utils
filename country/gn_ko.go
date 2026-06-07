//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Korean, "기니")
	dataGuinea.RegisterOfficialName(xlanguage.Korean, "기니 공화국")
	dataGuinea.RegisterCapital(xlanguage.Korean, "코나크리")
}
