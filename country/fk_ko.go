//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.Korean, "포클랜드 제도")
	dataFalklandIslands.RegisterOfficialName(xlanguage.Korean, "포클랜드 제도")
	dataFalklandIslands.RegisterCapital(xlanguage.Korean, "스탠리")
}
