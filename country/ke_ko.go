//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Korean, "케냐")
	dataKenya.RegisterOfficialName(xlanguage.Korean, "케냐 공화국")
	dataKenya.RegisterCapital(xlanguage.Korean, "나이로비")
}
