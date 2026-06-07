//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Php.RegisterName(xlanguage.Korean, "필리핀 페소")
}
