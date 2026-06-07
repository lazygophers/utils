//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Korean, "앵귈라")
	dataAnguilla.RegisterOfficialName(xlanguage.Korean, "앵귈라")
	dataAnguilla.RegisterCapital(xlanguage.Korean, "더밸리")
}
