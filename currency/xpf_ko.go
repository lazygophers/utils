//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xpf.RegisterName(xlanguage.Korean, "CFP 프랑")
}
