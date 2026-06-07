//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Irr.RegisterName(xlanguage.Korean, "이란 리알")
}
