//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Idr.RegisterName(xlanguage.Korean, "인도네시아 루피아")
}
