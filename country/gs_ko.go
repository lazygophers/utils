//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Korean, "사우스조지아 사우스샌드위치 제도")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Korean, "사우스조지아 사우스샌드위치 제도")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Korean, "킹에드워드포인트")
}
