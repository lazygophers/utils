//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gbp.RegisterName(xlanguage.Korean, "파운드 스털링")
}
