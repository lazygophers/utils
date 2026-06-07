//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Korean, "페로 제도")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Korean, "페로 제도")
	dataFaroeIslands.RegisterCapital(xlanguage.Korean, "토르스하운")
}
