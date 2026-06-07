//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SGD.RegisterName(xlanguage.Korean, "싱가포르 달러")
}
