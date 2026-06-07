//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zwl.RegisterName(xlanguage.Korean, "짐바브웨 달러")
}
