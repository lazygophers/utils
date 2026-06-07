//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lak.RegisterName(xlanguage.Korean, "라오스 킵")
}
