//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Korean, "과들루프")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Korean, "과들루프")
	dataGuadeloupe.RegisterCapital(xlanguage.Korean, "바스테르")
}
