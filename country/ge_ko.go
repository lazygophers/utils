//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Korean, "조지아")
	dataGeorgia.RegisterOfficialName(xlanguage.Korean, "조지아")
	dataGeorgia.RegisterCapital(xlanguage.Korean, "트빌리시")
}
