//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sar.RegisterName(xlanguage.Korean, "사우디아라비아 리얄")
}
