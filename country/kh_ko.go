//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Korean, "캄보디아")
	dataCambodia.RegisterOfficialName(xlanguage.Korean, "캄보디아 왕국")
	dataCambodia.RegisterCapital(xlanguage.Korean, "프놈펜")
}
