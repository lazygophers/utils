//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	RUB.RegisterName(xlanguage.Korean, "러시아 루블")
}
