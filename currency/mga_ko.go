//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mga.RegisterName(xlanguage.Korean, "아리아리")
}
