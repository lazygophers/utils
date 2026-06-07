//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lyd.RegisterName(xlanguage.Korean, "리비아 디나르")
}
