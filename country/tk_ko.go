//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Korean, "토켈라우")
	dataTokelau.RegisterOfficialName(xlanguage.Korean, "토켈라우")
}
