//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gip.RegisterName(xlanguage.Korean, "지브롤터 파운드")
}
