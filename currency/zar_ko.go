//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zar.RegisterName(xlanguage.Korean, "랜드")
}
