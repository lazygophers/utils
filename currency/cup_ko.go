//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cup.RegisterName(xlanguage.Korean, "쿠바 페소")
}
