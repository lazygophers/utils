//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bwp.RegisterName(xlanguage.Korean, "보츠와나 풀라")
}
