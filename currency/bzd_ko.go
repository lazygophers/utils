//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bzd.RegisterName(xlanguage.Korean, "벨리즈 달러")
}
