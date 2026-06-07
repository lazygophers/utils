//go:build lang_ko || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Hnl.RegisterName(xlanguage.Korean, "렘피라")
}
