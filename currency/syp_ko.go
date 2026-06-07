//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Syp.RegisterName(xlanguage.Korean, "시리아 파운드")
}
