//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uzs.RegisterName(xlanguage.Korean, "우즈베키스탄 숨")
}
