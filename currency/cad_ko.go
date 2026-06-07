//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cad.RegisterName(xlanguage.Korean, "캐나다 달러")
}
