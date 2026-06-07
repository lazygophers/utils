//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Korean, "지브롤터")
	dataGibraltar.RegisterOfficialName(xlanguage.Korean, "지브롤터")
	dataGibraltar.RegisterCapital(xlanguage.Korean, "지브롤터")
}
