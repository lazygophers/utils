//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Khr.RegisterName(xlanguage.Korean, "리엘")
}
