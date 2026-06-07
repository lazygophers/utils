//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mmk.RegisterName(xlanguage.Korean, "미얀마 짜트")
}
