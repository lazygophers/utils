//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Korean, "적도 기니")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Korean, "적도 기니 공화국")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Korean, "말라보")
}
