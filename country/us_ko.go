//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Korean, "미국")
	dataUnitedStates.RegisterOfficialName(xlanguage.Korean, "아메리카 합중국")
	dataUnitedStates.RegisterCapital(xlanguage.Korean, "워싱턴 D.C.")
}
