//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uah.RegisterName(xlanguage.Korean, "우크라이나 흐리우냐")
}
