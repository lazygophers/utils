//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gyd.RegisterName(xlanguage.Korean, "가이아나 달러")
}
