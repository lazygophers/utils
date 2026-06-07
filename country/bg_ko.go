//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Korean, "불가리아")
	dataBulgaria.RegisterOfficialName(xlanguage.Korean, "불가리아 공화국")
	dataBulgaria.RegisterCapital(xlanguage.Korean, "소피아")
}
