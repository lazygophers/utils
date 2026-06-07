//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Isk.RegisterName(xlanguage.Korean, "아이슬란드 크로나")
}
