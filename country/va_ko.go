//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Korean, "바티칸 시국")
	dataVaticanCity.RegisterOfficialName(xlanguage.Korean, "바티칸 시국")
	dataVaticanCity.RegisterCapital(xlanguage.Korean, "바티칸시")
}
