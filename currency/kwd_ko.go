//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kwd.RegisterName(xlanguage.Korean, "쿠웨이트 디나르")
}
