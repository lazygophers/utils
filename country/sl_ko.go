//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Korean, "시에라리온")
	dataSierraLeone.RegisterOfficialName(xlanguage.Korean, "시에라리온 공화국")
	dataSierraLeone.RegisterCapital(xlanguage.Korean, "프리타운")
}
