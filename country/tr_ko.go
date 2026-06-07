//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Korean, "튀르키예")
	dataTurkey.RegisterOfficialName(xlanguage.Korean, "튀르키예 공화국")
	dataTurkey.RegisterCapital(xlanguage.Korean, "앙카라")
}
