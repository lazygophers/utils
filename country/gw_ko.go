//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Korean, "기니비사우")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Korean, "기니비사우 공화국")
	dataGuineaBissau.RegisterCapital(xlanguage.Korean, "비사우")
}
