//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cve.RegisterName(xlanguage.Korean, "카보베르데 에스쿠도")
}
