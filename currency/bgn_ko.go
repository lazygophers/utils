//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bgn.RegisterName(xlanguage.Korean, "불가리아 레프")
}
