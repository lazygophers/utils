//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Korean, "콩고 민주 공화국")
	dataDrCongo.RegisterOfficialName(xlanguage.Korean, "콩고 민주 공화국")
	dataDrCongo.RegisterCapital(xlanguage.Korean, "킨샤사")
}
