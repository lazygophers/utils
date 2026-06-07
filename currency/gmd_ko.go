//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gmd.RegisterName(xlanguage.Korean, "달라시")
}
