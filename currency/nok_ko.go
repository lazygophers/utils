//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nok.RegisterName(xlanguage.Korean, "노르웨이 크로네")
}
