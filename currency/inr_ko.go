//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	INR.RegisterName(xlanguage.Korean, "인도 루피")
}
