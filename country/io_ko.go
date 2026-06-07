//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Korean, "영국령 인도양 지역")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Korean, "영국령 인도양 지역")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Korean, "디에고가르시아")
}
