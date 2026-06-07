//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	USD.RegisterName(xlanguage.Korean, "미국 달러")
}
