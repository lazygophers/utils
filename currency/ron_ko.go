//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ron.RegisterName(xlanguage.Korean, "루마니아 레우")
}
