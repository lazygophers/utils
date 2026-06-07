//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Korean, "에리트레아")
	dataEritrea.RegisterOfficialName(xlanguage.Korean, "에리트레아")
	dataEritrea.RegisterCapital(xlanguage.Korean, "아스마라")
}
