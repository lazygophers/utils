//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Korean, "세이셸")
	dataSeychelles.RegisterOfficialName(xlanguage.Korean, "세이셸 공화국")
	dataSeychelles.RegisterCapital(xlanguage.Korean, "빅토리아")
}
