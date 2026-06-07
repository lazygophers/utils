//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Byn.RegisterName(xlanguage.Korean, "벨라루스 루블")
}
