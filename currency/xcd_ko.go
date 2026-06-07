//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xcd.RegisterName(xlanguage.Korean, "동카리브 달러")
}
