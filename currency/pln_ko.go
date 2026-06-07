//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pln.RegisterName(xlanguage.Korean, "폴란드 즈워티")
}
