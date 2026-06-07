//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bnd.RegisterName(xlanguage.Korean, "브루나이 달러")
}
