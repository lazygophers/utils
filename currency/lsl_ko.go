//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lsl.RegisterName(xlanguage.Korean, "로티")
}
