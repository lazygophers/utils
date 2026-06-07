//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Jod.RegisterName(xlanguage.Korean, "요르단 디나르")
}
