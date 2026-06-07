//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Thb.RegisterName(xlanguage.Korean, "태국 바트")
}
