//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ils.RegisterName(xlanguage.Korean, "신 이스라엘 셰켈")
}
