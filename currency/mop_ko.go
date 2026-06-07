//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mop.RegisterName(xlanguage.Korean, "마카오 파타카")
}
