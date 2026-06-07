//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Vuv.RegisterName(xlanguage.Korean, "바누아투 바투")
}
