//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pab.RegisterName(xlanguage.Korean, "발보아")
}
