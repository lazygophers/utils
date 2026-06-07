//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Aed.RegisterName(xlanguage.Korean, "아랍에미리트 디르함")
}
