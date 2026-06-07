//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Korean, "니제르")
	dataNiger.RegisterOfficialName(xlanguage.Korean, "니제르 공화국")
	dataNiger.RegisterCapital(xlanguage.Korean, "니아메")
}
