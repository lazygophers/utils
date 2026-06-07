//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kzt.RegisterName(xlanguage.Korean, "카자흐스탄 텡게")
}
