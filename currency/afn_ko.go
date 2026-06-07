//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Afn.RegisterName(xlanguage.Korean, "아프가니")
}
