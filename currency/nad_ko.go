//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nad.RegisterName(xlanguage.Korean, "나미비아 달러")
}
