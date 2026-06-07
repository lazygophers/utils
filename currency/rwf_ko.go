//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Rwf.RegisterName(xlanguage.Korean, "르완다 프랑")
}
