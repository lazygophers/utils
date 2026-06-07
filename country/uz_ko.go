//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Korean, "우즈베키스탄")
	dataUzbekistan.RegisterOfficialName(xlanguage.Korean, "우즈베키스탄 공화국")
	dataUzbekistan.RegisterCapital(xlanguage.Korean, "타슈켄트")
}
