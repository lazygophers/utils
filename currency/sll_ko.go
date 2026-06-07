//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sll.RegisterName(xlanguage.Korean, "시에라리온 레온")
}
