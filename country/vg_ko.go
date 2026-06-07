//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Korean, "영국령 버진아일랜드")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Korean, "버진아일랜드")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Korean, "로드타운")
}
