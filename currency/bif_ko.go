//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bif.RegisterName(xlanguage.Korean, "부룬디 프랑")
}
