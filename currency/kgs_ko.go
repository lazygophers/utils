//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kgs.RegisterName(xlanguage.Korean, "키르기스스탄 솜")
}
