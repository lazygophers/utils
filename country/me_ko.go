//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Korean, "몬테네그로")
	dataMontenegro.RegisterOfficialName(xlanguage.Korean, "몬테네그로")
	dataMontenegro.RegisterCapital(xlanguage.Korean, "포드고리차")
}
