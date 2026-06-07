//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Korean, "르완다")
	dataRwanda.RegisterOfficialName(xlanguage.Korean, "르완다 공화국")
	dataRwanda.RegisterCapital(xlanguage.Korean, "키갈리")
}
