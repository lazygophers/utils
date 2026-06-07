//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Npr.RegisterName(xlanguage.Korean, "네팔 루피")
}
