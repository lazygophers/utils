//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Korean, "코코스 제도")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Korean, "코코스 제도")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Korean, "웨스트아일랜드")
}
