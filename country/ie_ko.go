//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Korean, "아일랜드")
	dataIreland.RegisterOfficialName(xlanguage.Korean, "아일랜드")
	dataIreland.RegisterCapital(xlanguage.Korean, "더블린")
}
