//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Korean, "그린란드")
	dataGreenland.RegisterOfficialName(xlanguage.Korean, "그린란드")
	dataGreenland.RegisterCapital(xlanguage.Korean, "누크")
}
