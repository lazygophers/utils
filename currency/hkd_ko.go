//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Hkd.RegisterName(xlanguage.Korean, "홍콩 달러")
}
