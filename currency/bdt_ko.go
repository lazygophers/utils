//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bdt.RegisterName(xlanguage.Korean, "방글라데시 타카")
}
