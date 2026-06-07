//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dop.RegisterName(xlanguage.Korean, "도미니카 페소")
}
