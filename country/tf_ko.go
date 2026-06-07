//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Korean, "프랑스령 남방 및 남극 지역")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Korean, "프랑스령 남방 및 남극 지역")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Korean, "포르토프랑세")
}
