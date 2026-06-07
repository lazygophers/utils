//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TWD.RegisterName(xlanguage.Korean, "신 타이완 달러")
}
