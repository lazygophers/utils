//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Aud.RegisterName(xlanguage.Korean, "오스트레일리아 달러")
}
