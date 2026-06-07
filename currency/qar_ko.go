//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Qar.RegisterName(xlanguage.Korean, "카타르 리얄")
}
