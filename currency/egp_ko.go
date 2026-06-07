//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Egp.RegisterName(xlanguage.Korean, "이집트 파운드")
}
