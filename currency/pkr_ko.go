//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pkr.RegisterName(xlanguage.Korean, "파키스탄 루피")
}
