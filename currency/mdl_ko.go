//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mdl.RegisterName(xlanguage.Korean, "몰도바 레우")
}
