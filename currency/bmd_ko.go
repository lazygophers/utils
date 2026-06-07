//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bmd.RegisterName(xlanguage.Korean, "버뮤다 달러")
}
