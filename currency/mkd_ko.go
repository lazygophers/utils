//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mkd.RegisterName(xlanguage.Korean, "마케도니아 데나르")
}
