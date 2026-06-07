//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Korean, "짐바브웨")
	dataZimbabwe.RegisterOfficialName(xlanguage.Korean, "짐바브웨 공화국")
	dataZimbabwe.RegisterCapital(xlanguage.Korean, "하라레")
}
