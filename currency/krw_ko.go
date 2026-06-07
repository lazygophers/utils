//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KRW.RegisterName(xlanguage.Korean, "대한민국 원")
}
