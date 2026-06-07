//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Korean, "영국")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Korean, "그레이트브리튼 북아일랜드 연합 왕국")
	dataUnitedKingdom.RegisterCapital(xlanguage.Korean, "런던")
}
