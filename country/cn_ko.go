//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Korean, "중국")
	dataChina.RegisterOfficialName(xlanguage.Korean, "중화인민공화국")
	dataChina.RegisterCapital(xlanguage.Korean, "베이징")
}
