//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kpw.RegisterName(xlanguage.Korean, "조선민주주의인민공화국 원")
}
