//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nzd.RegisterName(xlanguage.Korean, "뉴질랜드 달러")
}
