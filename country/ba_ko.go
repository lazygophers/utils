//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Korean, "보스니아 헤르체고비나")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Korean, "보스니아 헤르체고비나")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Korean, "사라예보")
}
