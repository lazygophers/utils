//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xof.RegisterName(xlanguage.Korean, "서아프리카 CFA 프랑")
}
