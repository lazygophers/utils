//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	JPY.RegisterName(xlanguage.Korean, "일본 엔")
}
