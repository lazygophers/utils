//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Amd.RegisterName(xlanguage.Korean, "아르메니아 드람")
}
