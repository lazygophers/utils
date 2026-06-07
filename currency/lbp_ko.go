//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lbp.RegisterName(xlanguage.Korean, "레바논 파운드")
}
