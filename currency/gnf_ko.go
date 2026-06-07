//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gnf.RegisterName(xlanguage.Korean, "기니 프랑")
}
