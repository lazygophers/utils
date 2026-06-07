//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Brl.RegisterName(xlanguage.Korean, "브라질 헤알")
}
