//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mvr.RegisterName(xlanguage.Korean, "몰디브 루피야")
}
