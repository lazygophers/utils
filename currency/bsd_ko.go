//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bsd.RegisterName(xlanguage.Korean, "바하마 달러")
}
